# taurshare
An introductory application to Redwood City.

# Setup
1. Create an [Autonomous Database](https://docs.oracle.com/en-us/iaas/autonomous-database/doc/adb.html).
    - Using Autonomous Database via Always Free is acceptable.
    - It would be beneficial to create a secondary, lesser-privileged user within "Database Actions".
2. Download the Oracle Wallet for your database by clicking "DB Connection".
    - "Instance wallet" is a preferable type.
    - It is recommended to use as long and as strong a password as possible.
3. Copy the TNS connection string for any service level you wish to experiment with.
4. Extract the downloaded ZIP file. Rename its directory to `wallet`.
    - Its original name may be similar to `Wallet_<database name>.zip`.
    - The resulting, renamed directory should contain `ewallet.p12` and `cwallet.sso`. Ensure these are present.
5. Place the `wallet` directory within the directory you are invoking `taurshare` from.
6. Copy `config.example.xml` to `config.xml` and edit appropriately.