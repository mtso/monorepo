const express = require('express')
const passport = require('passport')
const JwtStrategy = require('passport-jwt').Strategy
const app = module.exports = express()

const jwtFromRequest = function(req) {
  const head = req.get('Authorization')
  if (head) {
    return head.split(' ')[1]
  }
  return null
}

const users = [
  'wiggs',
  'mtso',
  'leaf',
]

const jwtStrategy = new JwtStrategy(
  {
    jwtFromRequest,
    secretOrKey: process.env.TOKEN_SECRET,
  },
  (payload, done) => {
    const username = payload.sub
    if (users.includes(username)) {
      done(null, {username})
    } else {
      done(null, false)
    }
  }
)

passport.serializeUser((user, cb) => cb(null, user))
passport.deserializeUser((obj, cb) => cb(null, obj))
passport.use(jwtStrategy)

app.use([
  passport.initialize(),
  passport.session(),
])

app.get('/test', passport.authenticate('jwt', {session: false}), (req, res) => {
  res.json({
    username: req.user.username,
    ok: true,
  })
})

app.listen(3750, () => console.log('listening on', 3750))
