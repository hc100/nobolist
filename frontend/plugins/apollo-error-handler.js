// make note it may be executed serverside
// https://nuxtjs.org/docs/2.x/internals-glossary/context
export default (
  { graphQLErrors, networkError, operation, forward },
  nuxtContext
) => {
  if (graphQLErrors) {
    graphQLErrors.map(({ message }) => {
      if (`${message}` === 'Access denied') {
        try {
          nuxtContext.app.context.redirect('/logout')
        } catch (e) {
          console.log(e)
        }
      } else {
        console.log(`[Error]: ${message}`) // eslint-disable-line
      }
    })
  }
  if (networkError) {
    console.log(`[Network error]: ${networkError}`) // eslint-disable-line
  }
}
