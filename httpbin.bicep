import radius as radius

@description('The Radius Application ID. Injected automatically by the rad CLI.')
param application string

resource gateway 'Applications.Core/gateways@2023-10-01-preview' = {
  name: 'gateway'
  properties: {
    application: application
    routes: [
      {
        path: '/'
        destination: 'http://httpbin:80'
      }
    ]
  }
}

resource httpbin 'Applications.Core/containers@2023-10-01-preview' = {
  name: 'httpbin'
  properties: {
    application: application
    container: {
      image: 'kennethreitz/httpbin:latest'
      ports: {
        api: {
          containerPort: 80
        }
      }
    }
  }
}
