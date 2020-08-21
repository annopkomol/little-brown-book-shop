import Vue from "vue"
import Vuetify from "vuetify/lib"

Vue.use(Vuetify)

export default new Vuetify({
  theme: {
    themes: {
      light: {
        primary: '#07B53B',
        accent: '#263147',
        secondary: '#FFFFFF',
        success: '#07B53B',
        info: '#C8C8C8',
        warning: '#263147',
        error: '#302F36'
      }
    },
  },
})
