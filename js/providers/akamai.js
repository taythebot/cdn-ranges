const bgpView = require('../utils/bgpview')

module.exports = () => {
  return ['12222', '20940'].map((asn) => {
    return bgpView(asn)
  })
}
