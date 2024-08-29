# techtest-multifinance-loan-sys
This is a simple loan engine application built using GoFiber with clean architecture principles. The app uses MySQL as the database, with GORM as the ORM.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Project Structure](#project-structure)
- [Database Schema](#database-schema)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Testing the Application](#testing-the-application)
- [API Endpoints](#api-endpoints)
    - [Create a Customer](#create-customer)
    - [Get Customer by ID](#get-customer-by-id)
    - [Retrieve customer details by NIK](#retrieve-customer-details-by-nik)
    - [Get Tenors by Customer ID](#get-tenors-by-customer-id)
    - [Update Lunas on Tenor](#update-islunas)
    - [Create Transaction](#create-transaction)
    - [Get Transactions by Customer ID](#get-transactions-by-customer-id)
- [Usage Example](#usage-example)
- [Contributing](#contributing)
- [License](#license)
- [Security Assessment](#security-assessment)
    - [Testing XSS](#testing-xss)
    - [Container Image Security assement (docker)](#container-image-security-assement)
    - [Testing for known security vulnerability](#testing-for-known-security-vulnerability)

## Prerequisites

- Go 1.18 or higher
- MySQL 5.7 or higher
- Docker (optional)

## Installation

1. **Clone the repository:**

```sh
$ git clone https://github.com/accalina/techtest-multifinance-loan-sys.git
$ cd techtest-multifinance-loan-sys
```
    
2. **Set up environment variables:** Create a `.env` file in the root directory and add MySQL connection details:

```dotenv
DB_DSN="root:fintech-password@tcp(mysql:3306)/loan_engine_db?charset=utf8mb4&parseTime=True&loc=Local"
```
    
3. **Install Go modules:**

```bash
$ go mod tidy
```

4. **Run the migrations:** The database schema will be automatically migrated when the application starts.
    

## Project Structure

```
loan-engine/
├── cmd/
│   └── main.go                   # Application entry point
├── config/
│   └── config.go                 # Configuration management
├── repository/
│   ├── customer_repository.go    # Customer repository interface and implementation
│   ├── tenor_repository.go       # Tenor repository interface and implementation
│   └── transaction_repository.go # Transaction detail repository interface and implementation
├── usecase/
│   ├── customer_usecase.go       # Customer business logic
│   ├── tenor_usecase.go          # Tenor business logic
│   └── transaction_usecase.go    # Transaction detail business logic
├── delivery/
│   ├── http/
│   │   ├── customer_handler.go    # Customer HTTP handler
│   │   ├── tenor_handler.go       # Tenor HTTP handler
│   │   └── transaction_handler.go # Transaction detail HTTP handler
├── entity/
│   ├── customer.go               # Customer entity/model
│   ├── tenor.go                  # Tenor entity/model
│   └── transaction_detail.go     # Transaction detail entity/model
├── infra/
│   ├── database.go               # Database connection setup
├── go.mod
└── go.sum
```

## Database Schema

The database schema consists of three tables:

1. **detail_customer**
    
    - `NIK` (Primary Key)
    - `FullName`
    - `LegalName`
    - `TempatLahir`
    - `TanggalLahir`
    - `Gaji`
    - `FotoKTP`
    - `FotoSelfie`
2. **tenor**
    
    - `ID` (Primary Key)
    - `CustomerID` (Foreign Key to `detail_customer`)
    - `Limit`
    - `MonthNumber`
3. **transaction_detail**
    
    - `ID` (Primary Key)
    - `CustomerID` (Foreign Key to `detail_customer`)
    - `OTRPrice` (On The Road Pricing)
    - `AdminFee`
    - `InstallmentAmount`
    - `InterestAmount`
    - `AssetName`

## Configuration

The configuration is managed using environment variables. The main configuration file is located in `config/config.go`. Ensure you have a `.env` file with the following content:

```dotenv
DB_DSN="root:fintech-password@tcp(mysql:3306)/loan_engine_db?charset=utf8mb4&parseTime=True&loc=Local"
```

## Running the Application

### Running via Docker
If you have docker and docker-compose installed you can run the image with the following command:

```bash
$ docker-compose up -d
```
the service will be build alongside mysql, so its ready to use on port 8080

### Running Manually
You can run the application with the following command:

```bash
$ go run cmd/main.go
```

or if you want to use Makefile config you can do with following command

```bash
$ make run
```

The application will start on port `8080` by default.

## Testing the Application

You can test the application with the following command:

```bash
$ go test mf-loan/delivery/http/tests mf-loan/repository/tests mf-loan/usecase/tests -v
```

or if you want to use Makefile config you can do with following command

```bash
$ make test
```

## API Endpoints

### Usage Example

### Customer
#### Create Customer

- **URL:** `/customers`
- **Method:** `POST`
- **Content-Type:** `application/json`
- **Body Example:**
```json
{
  "NIK": "1234567890123456",
  "FullName": "John Doe",
  "LegalName": "Johnathan Doe",
  "TempatLahir": "Jakarta",
  "TanggalLahir": "1990-01-01",
  "Gaji": 15000000,
  "FotoKTP": "/path/to/ktp.jpg",
  "FotoSelfie": "/path/to/selfie.jpg"
}
```

#### Get Customer by ID

- **URL:** `/customers/:id`
- **Method:** `GET`
- **URL Params:**
    - `id=[string]` (NIK of the customer)


1. **Create a customer:**
```bash
$ curl -X POST http://localhost:8080/customers \
-H "Content-Type: application/json" \
-d '{
    "nik": "1234567890123456",
    "full_name": "John Doe",
    "legal_name": "Johnathan Doe",
    "tempat_lahir": "Jakarta",
    "tanggal_lahir": "1990-01-01",
    "gaji": 10000000,
    "foto_ktp": "/path/to/ktp.jpg",
    "foto_selfie": "/path/to/selfie.jpg"
}'
```
    
2. **Retrieve customer details by NIK:**
```bash
$ curl -X GET http://localhost:8080/customers/1234567890123456
```
### Tenor
#### Create Tenor
- **URL:** `/tenors`
- **Method:** `POST`
- **Content-Type:** `application/json`
- **Body Example:**

```bash
$ curl -X POST http://localhost:8080/tenors \
-H "Content-Type: application/json" \
-d '{
    "customer_id": "1234567890123456",
    "limit": 1000000,
    "month_number": 1
}'
```

#### **Get Tenors by Customer ID**

- **URL:** `/customers/:customer_id/tenors`
- **Method:** `GET`
- **URL Params:**
    - `customer_id=[string]` (NIK of the customer)

#### **Update isLunas**

- **URL:** `/tenors/:id/lunas`
- **Method:** `PATCH`
```bash
$ curl -X PATCH http://localhost:8080/tenors/:id/lunas \
-H "Content-Type: application/json" \
-d '{
    "is_lunas": true
}'
```

### Transaction
#### **Create Transaction**

- **URL:** `/transactions`
- **Method:** `POST`
- **Content-Type:** `application/json`
- **Body Example:**

```bash
$ curl -X POST http://localhost:8080/transactions \
-H "Content-Type: application/json" \
-d '{
    "customer_id": "1234567890123456",
    "otr_price": 250000000,
    "admin_fee": 5000000,
    "installment_amount": 5000000,
    "interest_amount": 10000000,
    "asset_name": "Toyota Avanza"
}'
```

#### **Get Transactions by Customer ID**

- **URL:** `/customers/:customer_id/transactions`
- **Method:** `GET`
- **URL Params:**
    - `customer_id=[string]` (NIK of the customer)

## Contributing

Contributions are welcome! Please fork this repository and submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.


## Security Assessment

### Testing XSS
```lua
===========================================================================
Testing [XSS from URL]...
===========================================================================
===========================================================================
[*] Test: [ 1/1 ] <-> 2024-08-28 16:34:22.467665
===========================================================================

[+] Target: 

 [ http://192.168.50.161:8080 ]

---------------------------------------------

[!] Hashing: 

 [ bdc3505125c02acf6782cbc55c910984 ] : [ http://192.168.50.161:8080/customers/XSS/transactions ]

---------------------------------------------

[*] Trying: 

http://192.168.50.161:8080/customers/">bdc3505125c02acf6782cbc55c910984/transactions

---------------------------------------------

[+] Vulnerable(s): 

 [IE7.0|IE6.0|NS8.1-IE] [NS8.1-G|FF2.0] [O9.02]

---------------------------------------------

=============================================
[*] Injection(s) Results:
=============================================

 [NOT FOUND] -> [ bdc3505125c02acf6782cbc55c910984 ] : [ http://192.168.50.161:8080/customers/XSS/transactions ]

===========================================================================
[*] Final Results:
===========================================================================

- Injections: 1
- Failed: 1
- Successful: 0
- Accur: 0.0 %

===========================================================================
```

### Container Image Security assement
This scan will asess secirity on the docker container 
```lua
2024-08-30T03:14:37.841+0700    INFO    Vulnerability scanning is enabled
2024-08-30T03:14:37.841+0700    INFO    Secret scanning is enabled
2024-08-30T03:14:37.841+0700    INFO    If your scanning is slow, please try '--scanners vuln' to disable secret scanning
2024-08-30T03:14:38.043+0700    INFO    Detected OS: debian
2024-08-30T03:14:38.043+0700    INFO    Detecting Debian vulnerabilities...
2024-08-30T03:14:38.043+0700    INFO    Number of language-specific files: 0

accalina/fintech-loan:1.0 (debian 12.5)

Total: 0 (UNKNOWN: 0, LOW: 0, MEDIUM: 0, HIGH: 0, CRITICAL: 0)
```

### Testing for known security vulnerability
```lua
--------------- Timing report ---------------
  hostgroups: min 1, max 100000
  rtt-timeouts: init 1000, min 100, max 10000
  max-scan-delay: TCP 1000, UDP 1000, SCTP 1000
  parallelism: min 0, max 0
  max-retries: 10, host-timeout: 0
  min-rate: 0, max-rate: 0
---------------------------------------------
NSE: Using Lua 5.4.
NSE: Arguments from CLI: 
NSE: Loaded 45 scripts for scanning.
NSE: Script Pre-scanning.
NSE: Starting runlevel 1 (of 1) scan.
Initiating NSE at 15:24
Completed NSE at 15:24, 0.00s elapsed
Initiating ARP Ping Scan at 15:24
Scanning 192.168.50.161 [1 port]
Packet capture filter (device enp0s3): arp and arp[18:4] = 0x080027F9 and arp[22:2] = 0xE9DC
Completed ARP Ping Scan at 15:24, 0.10s elapsed (1 total hosts)
Overall sending rates: 10.46 packets / s, 439.52 bytes / s.
mass_rdns: Using DNS server 192.168.50.1
Initiating Parallel DNS resolution of 1 host. at 15:24
mass_rdns: 0.00s 0/1 [#: 1, OK: 0, NX: 0, DR: 0, SF: 0, TR: 1]
Completed Parallel DNS resolution of 1 host. at 15:24, 0.00s elapsed
DNS resolution of 1 IPs took 0.00s. Mode: Async [#: 1, OK: 1, NX: 0, DR: 0, SF: 0, TR: 1, CN: 0]
Initiating SYN Stealth Scan at 15:24
Scanning maverick (192.168.50.161) [1 port]
Packet capture filter (device enp0s3): dst host 192.168.50.110 and (icmp or icmp6 or ((tcp) and (src host 192.168.50.161)))
Discovered open port 8080/tcp on 192.168.50.161
Completed SYN Stealth Scan at 15:24, 0.01s elapsed (1 total ports)
Overall sending rates: 82.51 packets / s, 3630.36 bytes / s.
NSE: Script scanning 192.168.50.161.
NSE: Starting runlevel 1 (of 1) scan.
Initiating NSE at 15:24
NSE: Starting http-litespeed-sourcecode-download against 192.168.50.161:8080.
NSE: [http-litespeed-sourcecode-download 192.168.50.161:8080] Trying to download the source code of /index.php
NSE: Starting http-tplink-dir-traversal against 192.168.50.161:8080.
NSE: [http-tplink-dir-traversal 192.168.50.161:8080] HTTP GET /help/../../etc/shadow
NSE: Starting http-vuln-cve2009-3960 against 192.168.50.161:8080.
NSE: Starting http-majordomo2-dir-traversal against 192.168.50.161:8080.
NSE: [http-majordomo2-dir-traversal 192.168.50.161:8080] HTTP GET maverick/cgi-bin/mj_wwwusr?passw=&list=GLOBAL&user=&func=help&extra=/../../../../../../../../etc/passwd
NSE: Starting http-phpmyadmin-dir-traversal against 192.168.50.161:8080.
NSE: [http-phpmyadmin-dir-traversal 192.168.50.161:8080] HTTP POST maverick/phpMyAdmin-2.6.4-pl1/libraries/grab_globals.lib.php
NSE: [http-phpmyadmin-dir-traversal 192.168.50.161:8080] POST DATA usesubform[1]=1&usesubform[2]=1&subform[1][redirect]=../../../../../etc/passwd&subform[1][cXIb8O3]=1
NSE: Starting http-adobe-coldfusion-apsa1301 against 192.168.50.161:8080.
NSE: Starting http-vuln-cve2013-0156 against 192.168.50.161:8080.
NSE: Starting http-avaya-ipoffice-users against 192.168.50.161:8080.
NSE: Starting http-vuln-cve2013-7091 against 192.168.50.161:8080.
NSE: [http-vuln-cve2013-7091 192.168.50.161:8080] Trying to detect if the server is vulnerable
NSE: [http-vuln-cve2013-7091 192.168.50.161:8080] GET /zimbra/res/I18nMsg,AjxMsg,ZMsg,ZmMsg,AjxKeys,ZmKeys,ZdMsg,Ajx%20TemplateMsg.js.zgz?v=091214175450&skin=../../../../../../../../../dev/null%00
NSE: [http-vuln-cve2013-7091 192.168.50.161:8080] GET /zimbra/res/I18nMsg,AjxMsg,ZMsg,ZmMsg,AjxKeys,ZmKeys,ZdMsg,Ajx%20TemplateMsg.js.zgz?v=091214175450&skin=../../../../../../../../../etc/passwd%00
NSE: Starting http-vuln-cve2013-6786 against 192.168.50.161:8080.
NSE: Starting http-awstatstotals-exec against 192.168.50.161:8080.
NSE: Starting http-vuln-cve2014-3704 against 192.168.50.161:8080.
NSE: [http-vuln-cve2014-3704 192.168.50.161:8080] adding admin user (username: 'jsagzdokcc'; passwd: 'tdepegflbm')
NSE: Starting http-axis2-dir-traversal against 192.168.50.161:8080.
NSE: Starting http-dlink-backdoor against 192.168.50.161:8080.
NSE: Starting http-vuln-wnr1000-creds against 192.168.50.161:8080.
NSE: Starting http-coldfusion-subzero against 192.168.50.161:8080.
NSE: Starting http-vuln-cve2014-8877 against 192.168.50.161:8080.
NSE: Starting http-huawei-hg5xx-vuln against 192.168.50.161:8080.
NSE: Starting http-vuln-cve2012-1823 against 192.168.50.161:8080.
NSE: Starting http-shellshock against 192.168.50.161:8080.
NSE: [http-shellshock 192.168.50.161:8080] Sending '() { :;}; echo; echo -n iikrkzh; echo dwqmjim' in HTTP headers:User-Agent,Cookie and Referer
NSE: Finished http-litespeed-sourcecode-download against 192.168.50.161:8080.
NSE: Finished http-tplink-dir-traversal against 192.168.50.161:8080.
NSE: Finished http-majordomo2-dir-traversal against 192.168.50.161:8080.
NSE: Finished http-adobe-coldfusion-apsa1301 against 192.168.50.161:8080.
NSE: Finished http-dlink-backdoor against 192.168.50.161:8080.
NSE: [http-avaya-ipoffice-users 192.168.50.161:8080] Unexpected response returned for 404 check: 429 Too Many Requests
NSE: Finished http-vuln-cve2013-6786 against 192.168.50.161:8080.
NSE: [http-awstatstotals-exec 192.168.50.161:8080] This does not look like Awstats Totals. Quitting.
NSE: Finished http-awstatstotals-exec against 192.168.50.161:8080.
NSE: Finished http-vuln-cve2014-3704 against 192.168.50.161:8080.
NSE: [http-axis2-dir-traversal 192.168.50.161:8080] This does not look like an Apache Axis2 installation.
NSE: Finished http-axis2-dir-traversal against 192.168.50.161:8080.
NSE: [http-vuln-cve2014-8877 192.168.50.161:8080] Sending GET '//cmdownloads/?CMDsearch=".base64_decode("cmt4dXlhZnRzaGJremti")."' request
NSE: Finished http-vuln-cve2014-8877 against 192.168.50.161:8080.
NSE: [http-huawei-hg5xx-vuln 192.168.50.161:8080] Unexpected response returned for 404 check: 429 Too Many Requests
NSE: Finished http-vuln-cve2012-1823 against 192.168.50.161:8080.
NSE: Finished http-phpmyadmin-dir-traversal against 192.168.50.161:8080.
NSE: Finished http-vuln-cve2013-0156 against 192.168.50.161:8080.
NSE: Finished http-avaya-ipoffice-users against 192.168.50.161:8080.
NSE: [http-vuln-cve2013-7091 192.168.50.161:8080] The website seems to be not vulnerable to this attack.
NSE: Finished http-vuln-cve2013-7091 against 192.168.50.161:8080.
NSE: Finished http-huawei-hg5xx-vuln against 192.168.50.161:8080.
NSE: [http-vuln-wnr1000-creds 192.168.50.161:8080] Unable to obtain the id
NSE: Finished http-vuln-wnr1000-creds against 192.168.50.161:8080.
NSE: Finished http-shellshock against 192.168.50.161:8080.
NSE: Finished http-vuln-cve2009-3960 against 192.168.50.161:8080.
NSE: Finished http-coldfusion-subzero against 192.168.50.161:8080.
Completed NSE at 15:24, 0.28s elapsed
Nmap scan report for maverick (192.168.50.161)
Host is up, received arp-response (0.049s latency).
Scanned at 2024-08-28 15:24:13 UTC for 1s

PORT     STATE SERVICE    REASON
8080/tcp open  http-proxy syn-ack ttl 64
MAC Address: 00:00:00:00:00:00 (Intel Corporate)
Final times for host: srtt: 49457 rttvar: 55827  to: 272765

NSE: Script Post-scanning.
NSE: Starting runlevel 1 (of 1) scan.
Initiating NSE at 15:24
Completed NSE at 15:24, 0.00s elapsed
Read from /usr/bin/../share/nmap: nmap-mac-prefixes nmap-protocols nmap-services.
Nmap done: 1 IP address (1 host up) scanned in 0.57 seconds
           Raw packets sent: 2 (72B) | Rcvd: 3 (124B)
```