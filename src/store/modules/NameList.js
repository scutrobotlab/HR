import axios from "axios";

export default {
  namespaced: true,
  state: {
    list: [],
    loaded: false,
  },
  mutations: {},
  actions: {
    get({ state }) {
      return axios.get("/api/applicant/name-list").then((response) => {
        state.list = response.data.list;
        state.loaded = true;
      });
    },
  },
};
