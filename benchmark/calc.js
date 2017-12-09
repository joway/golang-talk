const helpers = require('../utils/helpers')

module.exports = function () {
  const beg = helpers.now()
  let sum = 0
  for (let i = 0; i < 1000000000; i++) {
    sum += 1000000000.123 * 1000000000.321 * 321.123456789
  }
  const end = helpers.now()

  helpers.log(`Nodejs : Calc used ${end - beg} ms`)
}
