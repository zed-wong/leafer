import colors from 'vuetify/es5/util/colors'
import i18n from './src/i18n'

export default {
  server: {
    host: '0',
    port: '3000',
  },
  // Disable server-side rendering: https://go.nuxtjs.dev/ssr-mode
  ssr: false,
  
  // Target: https://go.nuxtjs.dev/config-target
  target: 'static',

  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    titleTemplate: 'Leafer',
    title: 'Leafer',
    htmlAttrs: {
      lang: 'en'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no, viewport-fit=cover' },
      { hid: 'description', name: 'description', content: '' },
      { name: 'format-detection', content: 'telephone=no' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  env: {
    client_id: '51186d7e-d488-417d-a031-b4e34f4fdf86',
    //pusdid: '965e5c6e-434c-3fa9-b780-c50f43cd955c',
    pusdid:'31d2ea9c-95eb-3355-b65b-ba096853bc18',
    developer_id: '44d9717d-8cae-4004-98a1-f9ad544dcfb1',
    tg_bot_id: 'leafer_telegram_bot',
    group_link: 'https://mixin.one/codes/ae820758-d13b-44f9-bcd8-6f43cea26c15',
    oauth_url: 'http://127.0.0.1/oauth/',
  },

  // Global CSS: https://go.nuxtjs.dev/config-css
  css: [
    '~/assets/global.css'
  ],

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    '~/plugins/global.js'
  ],

  // Auto import components: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    '@nuxt/typescript-build',
    // https://go.nuxtjs.dev/vuetify
    '@nuxtjs/vuetify',
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    [
      '@nuxtjs/i18n',
      {
        locales: [
          {
            code: "en",
            iso: "en-US",
            file: "en-US.json"
          },
          {
            code: "zh",
            iso: "zh-TW",
            file: "zh-TW.json"
          }
        ],
        langDir: 'lang/',
        defaultLocale: "en",
        vueI18n: i18n,
      }
    ]
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    // Workaround to avoid enforcing hard-coded localhost:3000: https://github.com/nuxt-community/axios-module/issues/308
    baseURL: '/',
    proxy: false
  },
  // proxy: {
  //   "/price":"http://127.0.0.1:8080",
  //   "/mixinoauth": "http://127.0.0.1:8080",
  //   "/poll":"http://127.0.0.1:8080",
  //   "/delete/vault": "http://127.0.0.1:8080",
  //   "/update/tg": "http://127.0.0.1:8080",
  //   "/update/ratio": "http://127.0.0.1:8080",
  //   "/update/number": "http://127.0.0.1:8080",
  // },

  // Vuetify module configuration: https://go.nuxtjs.dev/config-vuetify
  vuetify: {
    customVariables: ['~/assets/variables.scss'],
    font: {
      family: '-apple-system,BlinkMacSystemFont,Inter,Roboto,Oxygen,Helvetica Neue,PingFang SC,sans-serif'
    },
    theme: {
      dark: false,
      themes: {
        dark: {
          primary: colors.blue.darken2,
          accent: colors.grey.darken3,
          secondary: colors.amber.darken3,
          info: colors.teal.lighten1,
          warning: colors.amber.base,
          error: colors.deepOrange.accent4,
          success: colors.green.accent3
        }
      }
    }
  },

  // Build Configuration: https://go.nuxtjs.dev/config-build
  build: {
    babel: {
      plugins: [
        [
          'component', {
            libraryName: 'maz-ui',
            styleLibraryName: 'css'
          }
        ]
      ]
    }
  },
  pwa: {
    workbox: {
      enabled: false,
      runtimeCaching: [
        {
          urlPattern:
            /^(https:\/\/images\.mixin\.one\/|https:\/\/mixin-images\.zeromesh\.net\/).*$/,
          handler: "CacheFirst",
        },
      ]
    },
  },
}
