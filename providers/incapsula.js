const axios = require('axios')

module.exports = async () => {
  const { data } = await axios.post(
    'https://my.incapsula.com/api/integration/v1/ips',
    'resp_format=text'
  )

  return data
    .split('\n')
    .slice(0, -1)
    .filter((e) => !e.includes(':'))
}
