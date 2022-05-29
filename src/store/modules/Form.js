import axios from "axios";

export default {
  namespaced: true,
  state: {
    form: null,
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
        state.form = response.data.form;
      });
    },
  },
};
