const axios = require('axios')

module.exports = async () => {
  const { data } = await axios.get(
    'https://cachefly.cachefly.net/ips/rproxy.txt'
  )

  return data.split('\n').slice(0, -1)
}
