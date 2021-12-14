# Nuri: Getting Blockchain Data

Api for getting blockchain block and transaction details

## Things you need

* [Go](https://golang.org/dl/): `brew install go`
* Install docker
* Setup `go` folder in your home directory: `/Users/<user_id>/go/src/github.com/Shahjahan-Miah-coding-challenge`. So, `go env` should show `GOPATH="/Users/<user_id>/go"`
* Clone `NuriCareers/Shahjahan-Miah-coding-challenge` repository in `/Users/<user_id>/go/src/github.com/Shahjahan-Miah-coding-challenge`: `git clone <repo_url>`
* After cloning, the `go.mod` file should be found in this directory `/Users/<user_id>/go/src/github.com/nuri/Shahjahan-Miah-coding-challenge/go.mod`
* Add `/Users/<user_id>/go/bin` in environment variable `PATH` if it is not already there.

## Running the App 

### Run Scripts
We use [Go Modules](https://blog.golang.org/using-go-modules) for dependency management.

### MAKEFILE IS YOUR FRIEND, ANY COMMAND YOU FIND BELOW IS ALREADY THERE , WITH HELP
1. `go mod vendor` should create a `vendor` directory in project root directory. (`/Users/<user_id>/go/src/github.com/nuri/Shahjahan-Miah-coding-challenge/vendor`)
If for any reason, this directory is not created then try to clear the cache and run the command again. To clear the cache run below command:
   
```shell
$ go clean -modcache  #clean module cache
$ go mod vendor       #setup vendor dir again
```
### Run docker compose

Start the **docker engine**. Next, run below command.
```shell
$ docker-compose up
```

### Run application

After docker run, make sure your api is running inside container and ready for connections...
```shell
$ go build .\cmd\server\main.go
$ go run .\cmd\server\main.go
```

Now application is listing port :8000 and ready to go.

### Transaction Details
###### http://localhost:8000/api/transaction/BTC/ee475443f1fbfff84ffba43ba092a70d291df233bd1428f3d09f7bd1a6054a1f
```
{
    "data": {
        "TXID": "ee475443f1fbfff84ffba43ba092a70d291df233bd1428f3d09f7bd1a6054a1f",
        "Fee": "0.0",
        "Time": "22/09/2012 10:45",
        "Sent_Value": "102.12000000"
    },
    "message": "SUCCESS!",
    "status": true
}
```

### Block Details
###### http://localhost:8000/api/block/BTC/000000000000034a7dedef4a161fa058a2d67a173a90155f3a2fe6fc132e0ebf
```shell
{
    "data": {
        "Network": "BTC",
        "BlockNo": 200000,
        "Time": "22/09/2012 10:45",
        "Previous_Block_Hash": "00000000000003a20def7a05a77361b9657ff954b2f2080e135ea6f5970da215",
        "Next_Block_Hash": "00000000000002e3269b8a00caf315115297c626f954770e8398470d7f387e1c",
        "Size": 247533,
        "Transactions": [
            {
                "TXID": "dbaf14e1c476e76ea05a8b71921a46d6b06f0a950f17c5f9f1a03b8fae467f10",
                "Fee": "0.0",
                "Time": "22/09/2012 10:45",
                "Sent_Value": "50.63517500"
            },
            {
                "TXID": "ee475443f1fbfff84ffba43ba092a70d291df233bd1428f3d09f7bd1a6054a1f",
                "Fee": "0.0",
                "Time": "22/09/2012 10:45",
                "Sent_Value": "102.12000000"
            },
            {
                "TXID": "e03a9a4b5c557f6ee3400a29ff1475d1df73e9cddb48c2391abdc391d8c1504a",
                "Fee": "0.0",
                "Time": "22/09/2012 10:45",
                "Sent_Value": "95.00000000"
            },
            {
                "TXID": "ffa0267c8f2af736858894d6f3e5081a05e2ec16dc98f78a80f376ce35077491",
                "Fee": "0.0",
                "Time": "22/09/2012 10:45",
                "Sent_Value": "373.25663465"
            },
            {
                "TXID": "2674c8a46b75e5a5e3287a34d7999a1e5e6f052be36f63f3cb483ec148c9b86c",
                "Fee": "0.00050000",
                "Time": "22/09/2012 10:45",
                "Sent_Value": "9466.22517855"
            },
            {
                "TXID": "6c3c876a1711a6851d8fca2c821415efd512f59d0ef32dffb3bb4dc90b7d4fa2",
                "Fee": "0.0",
                "Time": "22/09/2012 10:45",
                "Sent_Value": "12.55343673"
            },
            {
                "TXID": "f7a0f2b9dfa73dc795ca42fc4104adaa0f8c8a268dbb9348c987ff037fcc73ae",
                "Fee": "0.00050000",
                "Time": "22/09/2012 10:45",
                "Sent_Value": "28238.25441550"
            },
            {
                "TXID": "54dc90aa618ea1c300aac021399c66f5f5152848a57984a757075036e3046147",
                "Fee": "0.0",
                "Time": "22/09/2012 10:45",
                "Sent_Value": "63.19597494"
            },
            {
                "TXID": "7fb40b4c18ca267aa5ab3a042ef8608c1f90e73ff171ce876c2d2242ed79c65d",
                "Fee": "0.00050000",
                "Time": "22/09/2012 10:45",
                "Sent_Value": "1.01342540"
            },
            {
                "TXID": "80efe43cf64a524d1417546a027786127ad87475f3af1c13b8f3719cd4268679",
                "Fee": "0.0",
                "Time": "22/09/2012 10:45",
                "Sent_Value": "50.00000000"
            }
        ]
    },
    "message": "Success",
    "status": true
}
```
### Logging (Optional)
To see the log level wise log, please specify the log level to the configuration.json
```
{
    "logLevel": "INFO"
}
```
### Authorization and Middleware (Optional)
JWT token authentication is implemented to auth\token.go file but not used. In the future, this api can use if needed. 
```
{
    "ApiSecret":  "It should be a strong key",
    "IsJwtEnabled": false
}
```
Middleware could be setup like below - 
```
$ router.HandleFunc("/api/block/{network}/{blockhash}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(controllers.GetBlockData))).Methods("GET")
$ router.HandleFunc("/api/transaction/{network}/{txid}", middleware.SetMiddlewareJSON(middleware.SetMiddlewareAuthentication(controllers.GetTransactionData))).Methods("GET")
```



### Dynamic Configuration loading (Optional)
Configuration.json file could be load in start up and get json parameter and set to the application if needed. Under internal\config file. Possible code snippets could be - 
```
if err := config.LoadJSONConfig(config.Config); err != nil {
	logging.Fatal(logTag, "unable to load configuration. error=%v", err)
}
logging.Info(logTag, "configuration file loaded")
logging.SetLogLevel(config.Config.LogLevel)
```

### Running K8s Locally With Kind (Optional)
Let’s start by installing this:
```
$ brew install kind
```
And then once this has been installed successfully, let’s create a new cluster:
```
$ kind create cluster
```
This will create a cluster that will just be called kind on our local machine. With this created, you’ll then be able to run:
```
$ kubectl cluster-info --context kind-kind
```
Awesome, we now have a cluster setup and ready to use and deploy to!

### Defining our K8s Deployment (Optional) 
Now that we have a kubernetes cluster up and running on our local machine, we can now look at creating a deployment.yaml for our blockchain API.The deployment.yml will effectively contain everything that Kubernetes needs to know in order to run our application within the cluster as a number of pods.Let’s dive in and start defining this. We’ll start off by creating a new directly under which this configuration will live called config and then we’ll create a new deployment.yml file within this:
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: blockchain.api
spec:
  replicas: 3
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 0
  selector:
    matchLabels:
      name: blockchain.api
  template:
    metadata:
      labels:
        name: blockchain.api
    spec:
      containers:
        - name: application
          image: "hasanshahjahan/blockchain-api:latest"
          imagePullPolicy: Always
          ports:
            - containerPort: 8000
```

Now we are ready to deploy our application to K8s.