export default (context) => {
  return {
    httpEndpoint: process.env.graphqlEndpoint,
    httpLinkOptions: {
      fetchOptions: {
        mode: 'cors',
      },
      credentials: 'omit',
    },
  }
}
