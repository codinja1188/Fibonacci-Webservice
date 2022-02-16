# Fibonacci Web Service
**Challenge:**

```Create a web service that accepts a number, n, as input and returns the first n Fibonacci numbers, starting from 0```

1. The project should provide a GoLang web service.

  a. The web service accepts a number, n, as input and returns the first n Fibonacci numbers, starting from 0.

  I.E. given n  = 5, appropriate output would represent the sequence [0, 1, 1, 2, 3].

  b. Given a negative number or non integer, respond with an appropriate error.

2. Include whatever instructions are necessary to build and deploy/run the project, where "deploy/run" means the web service is accepting requests and responding to them as appropriate.

Issued from the location where you downloaded the project. The service will be
available on port ``443`` of your host operating system, eg::

  ``https://localhost/fbseries``

NOTE:
=====
'''
    command to create a self signed certificate

    $ openssl req -subj '/CN=myplatform.host' -x509 -newkey rsa:4096 -nodes -keyout server.key -out server.crt -days 36500

    we have enable Chrome settings to accept self-signed localhost certificate

    => Simply paste this in your chrome:

        $ chrome://flags/#allow-insecure-localhost

    => You should see highlighted text saying: Allow invalid certificates for resources loaded from localhost

        $ Click Enable.
'''    

Thoughts about a production web service:
```
1) Security
      * encryption 
        Here we are added encryption, the server is started with our self signed certificates.
2) Performance
      * calculate the full sequence every time, or store sequences that are already called?
      * Right now, n is not limited. A user can enter large number which can cause memory issues
3) Scalability-
      * Load balancer
      * Ability to scale out number of web servers
4) High availability
      * Handled by application layer or VM infrastructure?
5) Monitoring / Alerting in case web service is unavailable
```
