# CDN Ranges
This is a tool to download a list of IP ranges used by CDNs (Cloudflare, Akamai, Incapsula, Fastly, etc). This helps to avoid performing unnecessary port scans when doing bug bounties.

This uses publicly available lists of IP ranges, provided by most providers, and [BGPView](https://bgpview.io/) to query IP ranges for ASNs.

This was heavily inspired by [Project Discovery's cdncheck](https://github.com/projectdiscovery/cdncheck).

## CDN Providers
* Akamai
* CacheFly
* CDNetworks
* Cloudflare
* CloudFront
* Fastly
* Incapsula
* Limelight
* MaxCDN
* StackPath
* Sucuri

If a provider is missing, please open an issue with a link to their IP ranges or ASN

## Usage
Download ip ranges for all providers
```
node download --output ranges.txt
```

Download for a specific provider (lowercase)
```
node download --provider cloudflare
```

Dump in json format
```
node download --format json --output ranges.json
```

Dump in csv format (provider,range)
```
node download --format csv --output ranges.csv
```

## Support Formats
* txt (default)
* json
* csv