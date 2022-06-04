import axios from "axios";

export default {
  namespaced: true,
  state: {
    fields: null,
    intent: null,
    loaded: false,
  },
  mutations: {},
  actions: {
    get({ state }) {
      if (state.loaded) return;
      return axios({
        method: "get",
        url: "/api/public/form",
      }).then((response) => {
        state.fields = response.data.fields;
        state.intent = response.data.intent;
      });
    },
  },
};
