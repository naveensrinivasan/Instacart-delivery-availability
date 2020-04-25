# Instacart delivery availability

This code checks for instacart delivery availability and notifies the availability via `iMessage`

This code uses the instacart API and doesn't do screen scraping.

## Settings

Update the settings.yaml 

```yaml
---
email: instacart-email address
password: instacart password
stores:    # stores to look for delivery availability
  - costco
  - harris-teeter
notification: 12334567890
```