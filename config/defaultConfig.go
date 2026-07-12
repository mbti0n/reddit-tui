package config

const defaultConfiguration = `
#
# Default configuration for reddittui.
# Uncomment to configure
#

#[core]
#bypassCache = false
#logLevel = "Warn"

#[filter]
#keywords = ["drama"]
#subreddits = ["news", "politics"]

#[client]
#timeoutSeconds = 10
#cacheTtlSeconds = 3600

#[server]
#domain = "old.reddit.com"
#type = "old"

# Color configuration
[colors]
accent = "#FF713E"
link = "#6688E4"
negative = "#6B5DFB"
text = "#F5F5F5"
subtext = "#D0D0D0"
`
