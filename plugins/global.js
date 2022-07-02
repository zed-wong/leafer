import Vue from 'vue';
import Bridge from '@foxone/mixin-sdk-jsbridge';
const bridge = new Bridge({
  client_id: process.env.client_id,
  logLevel: 'warn',
})

//this is to help Webstorm with autocomplete
Vue.prototype.$bridge = bridge;

export default ({ app }, inject) => {
  inject('bridge', bridge);

}
