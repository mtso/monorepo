const articles = [
  'A',
  'The',
]

const nouns = [
  'behest', 
  'carport', 
  'crocodile', 
  'legume', 
  'miter', 
  'ping', 
  'shearling', 
  'streetcar', 
  'terrapin', 
  'tinderbox',
  'chaise',
  'control', 
  'covariate', 
  'cynic', 
  'dock', 
  'freedom', 
  'locomotive', 
  'pantyhose', 
  'radiosonde', 
  'stamp',
]

const verbs = [
  'invents',
  'whispers',
  'scares',
  'troubles',
  'bangs',
  'turns',
  'boasts',
  'sins',
  'deceives',
  'forces',
  'launches',
  'preserves',
]

const images = [
  'https://images.unsplash.com/photo-1428591850870-56971c19c3d9?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=fecc1c91f371ba8afc7d216b75d21445',
  'https://images.unsplash.com/photo-1489715100843-d2e89ba2d60f?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=a7407e62fb1c656bf9b7a5fe539468b2',
  'https://images.unsplash.com/photo-1470138000694-6580a25339f7?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=747a26c797492f52183df19ab392e47a',
  'https://images.unsplash.com/photo-1489258492820-d7b57603ec9e?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=5d4f5b1353dd9cdc56a19a6ce9288cb6',
  'https://images.unsplash.com/photo-1473483376464-ab7087f1d07d?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=3852a9ec7154bfef50b0f5f904a5dbd5',
  'https://images.unsplash.com/photo-1497444142752-f88f16400734?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=8ddd3ce968e0ac28c37c6c581072ba3a',
  'https://images.unsplash.com/photo-1484506730629-29bfeee38346?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=c2a793da8e12e6c0818e5c20079933b2',
  'https://images.unsplash.com/photo-1465352320358-78e22b587ac6?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=14fe076ba89a20c1a82f4fb24d06d845',
  'https://images.unsplash.com/photo-1437448317784-3a480be9571e?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=a9b0f506e313a033dac80dbebcc40253',
  'https://images.unsplash.com/photo-1453813063438-1ed846c7daf0?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=79e4e7bd6b20efecfa9d8f635bcfabb4',
  'https://images.unsplash.com/photo-1474729036944-a6eb9d43e14d?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=26717888428ec4bd5ce2ad67edec8976',
  'https://images.unsplash.com/photo-1466265019454-e865c0399ca4?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=a59cb1908f9b920e519dba96316b87ba',
  'https://images.unsplash.com/photo-1489547821686-3512bfb79f91?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=de91122bf4a4d4356b16b169748fd7d5',
  'https://images.unsplash.com/photo-1491425432462-010715fd7ed7?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=2db3d00def91d2d01dfa730d1165635d',
  'https://images.unsplash.com/photo-1464440579920-df688e377d3c?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=dc7c5d43404530d787ded74e89457de2',
  'https://images.unsplash.com/photo-1498183249004-bf61ac7295f2?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=dd6553b83bf35c1e44cf124d7c0bcb9a',
  'https://images.unsplash.com/photo-1468850726958-63df2c33e0ca?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=a1b52c49e2a644982f6d4592c4ec6e90',
  'https://images.unsplash.com/photo-1474804260312-e2e2a705f7f9?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=e141e936c3739ecfd55bb0193c954196',
  'https://images.unsplash.com/photo-1494301950624-2c54cc9826c5?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=c5335c52a745b95392f4ba667cc47377',
  'https://images.unsplash.com/photo-1462331940025-496dfbfc7564?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=35158310bf85902b98ea814f11e81d5b',
  'https://images.unsplash.com/photo-1444208393177-b2a88904ed8d?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=8e7afd3e42b630325aa2197507f71418',
  'https://images.unsplash.com/photo-1497293194408-f99df0ef85bf?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=26d9a049482d4cf6cabf54761ab5e845',
  'https://images.unsplash.com/photo-1496436818536-e239445d3327?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=893bbce24ae9995507f6658b85f0a3df',
  'https://images.unsplash.com/photo-1453329180519-b4dba097ed5b?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=c7165931ad8fae9144526ba119a9450d',
  'https://images.unsplash.com/photo-1488235742400-36898425c618?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=2016d0253100b5ad5d692206db768dbc',
  'https://images.unsplash.com/photo-1498457134284-b5317c652a1e?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=d56e7c6600b936463cedd2bcaf6497bf',
  'https://images.unsplash.com/photo-1497535944603-98de35a7eef9?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=6b1285dac20da105b853c77a9a4aa273',
  'https://images.unsplash.com/photo-1476642514879-ef4820d8c992?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=2666518386f274899edfd6457583f7f0',
  'https://images.unsplash.com/photo-1485148470689-9e57a229c0f8?ixlib=rb-0.3.5&q=80&fm=jpg&crop=entropy&cs=tinysrgb&w=300&h=400&fit=crop&s=0fbaba4a770f67971612394c14e0fc88',
]

const randInt = (max) => {
  return Math.floor(Math.random() * max)
}

class Generator {
  iter = 0

  createCaption = () => {
    return [
      articles[randInt(2)],
      nouns[randInt(nouns.length)],
      verbs[randInt(verbs.length)],
      articles[randInt(2)].toLowerCase(),
      nouns[randInt(nouns.length)],
    ].join(' ') + '.'
  }

  getFive = () => {
    const pics = []

    for (let i = 0; i < 5; i++) {
      pics.unshift(this.createPic())
    }

    return pics
  }

  createPic = () => {
    const id = ++this.iter
    const caption = this.createCaption()
    const image_url = images[randInt(images.length)]

    return {
      id,
      caption,
      image_url,
    }
  }

  getInitialPics = () => {
    const pics = []

    for (let i = 0; i < 10; i++) {
      pics.unshift(this.createPic())
    }

    return pics
  }
}

export {
  Generator,
}
