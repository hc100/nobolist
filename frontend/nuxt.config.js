export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    title: 'nobolist',
    htmlAttrs: {
      lang: 'ja'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/eslint
    '@nuxtjs/eslint-module',
    // https://go.nuxtjs.dev/tailwindcss
    '@nuxtjs/tailwindcss',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    '@nuxtjs/apollo',
  ],

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
  },

  apollo: {
    clientConfigs: {
      default: '~/plugins/apollo-config.js',
    },
    defaultOptions: {
      // See 'apollo' definition
      // For example: default query options
      // https://www.apollographql.com/docs/react/api/core/ApolloClient/#apolloclient-functions
      $query: {
        loadingKey: 'loading',
        fetchPolicy: 'cache-first', // https://www.apollographql.com/docs/react/data/queries/#setting-a-fetch-policy
      },
    },
    watchLoading: '~/plugins/apollo-watch-loading-handler.js',
    errorHandler: '~/plugins/apollo-error-handler.js',
  },

  env: {
    graphqlEndpoint: process.env.GRAPHQL_ENDPOINT || 'http://localhost:8081/query',
  },
}
