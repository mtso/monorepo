const crypto = require('crypto')
const express = require('express')
const session = require('express-session')
const bodyParser = require('body-parser')
const mongoose = require('mongoose')
const passport = require('passport')
const LocalStrategy = require('passport-local').Strategy

/*
 * Set Up Mongoose
 */
const User = mongoose.model('User', new mongoose.Schema({
  username: {type: String, unique: true},
  email: {type: String, unique: true},
  password: String,
  verified_at: {type: Date, default: null},
  verification_token: String,
}))

mongoose.connect(process.env.MONGODB_URI)

/*
 * Set Up Passport
 */
passport.serializeUser(function(user, done) {
  done(null, user.id)
})

// passport.deserializeUser(function(id, done) {
//   User.findById(id, done)
// })
passport.deserializeUser(User.findById.bind(User)) // :O

passport.use(new LocalStrategy((username, password, done) => {
  User.findOne({ username }, function(err, user) {
    if (err) { return done(err) }

    if (!user) { return done(null, false, {message: 'Incorrect username'}) }

    if (user.password !== password) { return done(null, false, {message: 'Incorrect password'}) }

    done(null, user)
  })
}))

/*
 * Set Up App-Level Middleware
 */
const app = module.exports = express()

app.use(session({
  secret: 'cupcakes',
  saveUninitialized: true,
  resave: true,
}))
app.use(bodyParser.json())
app.use(passport.initialize())
app.use(passport.session())

/*
 * Set Up Controllers
 */
app.get('/', (req, res) => {
  var username = (req.user && req.user.username) || ''
  res.status(200).send(`hi~ ${username}`)
})

app.get('/signup', attachToken, (req, res) => {
  var query = req.query
  var username = query.username
  var password = query.password
  var email = query.email
  var verification_token = query.token

  if (!username || !password || !email) {
    return res.json({
      ok: false,
      error: 'Missing username, password, or email',
    })
  }

  User.create({
    username,
    password,
    email,
    verification_token,
  }, (err, user) => {
    if (err) {
      if (err.code === 11000) {
        res.json({ok: false, error: 'Username or email exists'})
      } else {
        res.json({ok: false, error: err})
      }
      return
    }
    res.json({
      ok: true,
      user,
    })
  })
})

app.get('/verify', (req, res) => {
  var verification_token = req.query.token

  User.findOneAndUpdate(
    { verification_token },
    {
      $unset: {verification_token: 1},
      verified_at: new Date(),
    },
    { new: true },
    (err, user) => {
      if (err) {
        return res.json({ ok: false, error: 'Server error' })
      }

      res.send(200, 'Verified')
    })
})

app.get('/login', passport.authenticate('local', {
  failureRedirect: '/invalid-login',
  successRedirect: '/',
}))

app.get('/invalid-login', (req, res) => { res.status(200).send('Invalid login') })

app.get('/logout', (req, res) => {
  req.logout()
  res.redirect('/')
})

// Protected resource
app.get('/chocolate', isAuthenticated, (req, res) => res.json({ ok: true, chocolate: getChocolate() }))

/*
 * Start App
 */
;((port) => app.listen(port, () => console.log('Listening on', port)))(process.env.PORT || '3750')

/*
 * Utility Middlewares
 */
function attachToken(req, res, next) {
  generateToken((err, token) => {
    if (err) {
      return res.json({
        ok: false,
        message: 'Server error',
      })
    }
    req.query.token = token
    next()
  })
}

function generateToken(cb) {
  crypto.randomBytes(64, (err, buf) => {
    cb(err, buf.toString('hex'))
  })
}

function isAuthenticated(req, res, next) {
  if (!req.user) {
    return res.redirect('/invalid-login')
  }
  next()
}

const chocolateBars = [
  'Hershey\'s Milk Chocolate',
  'Ghirardelli',
  'Mast Brothers'
]

function getChocolate() {
  return chocolateBars[Math.floor(Math.random() * chocolateBars.length)]
}
