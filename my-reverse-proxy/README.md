# Reverse Proxy in Go

This project implements a simple reverse proxy in Go. It forwards incoming HTTP requests to a target server and manages the responses back to the client.

In micro service landscape , to run single application we need to configure various urls routes to separate deployments environments. Yes, I could achive by Nginx but I don't wanted teach every developer to support Nginx. So written this simple reverse proxy, it served my purpose. 

## Featues
   - Customize target server depends on app context



### Running the Reverse Proxy

To run the reverse proxy, execute the following command in the `root` directory:

```
go run main.go
```

By default, the proxy will listen on port 8083. You can change the port and target server in the `main.go` file.

### Configuration

You can configure the target server and other options directly in the `config.yml` file. Make sure to update the target URL to the server you want to proxy requests to.

### Example Usage

Once the server is running, you can send requests to the reverse proxy:

```
curl http://localhost:8083/
```

The proxy will forward the request to the configured target server and return the response.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.