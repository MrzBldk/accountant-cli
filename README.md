# Accountant CLI
A CLI for managing user accounts
# Guide
Command for creating a user: ```accountant-cli user create <username> --balance=<balance>```

Command for displaying user info: ```accountant-cli user <username>```

Command for debeting a user: ```accountant-cli credit <username> --amount=<amount> --description==<description>```

Command for crediting a user: ```accountant-cli debit <username> --amount=<amount> --description==<description>```
# Used libraries
- Cobra
- GORM
