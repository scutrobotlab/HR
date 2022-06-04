import axios from "axios";

export default {
  namespaced: true,
  state: {
    done: null,
    form_end: null,
    form_start: null,
    time_end: null,
    time_start: null,
  },
  mutations: {},
  actions: {
    get({ state }) {
      return axios({
        method: "get",
        url: "/api/public/time-frame"
      }).then((response) => {
        state.done = response.data.done;
        state.form_end = response.data.form_end;
        state.form_start = response.data.form_start;
        state.time_end = response.data.time_end;
        state.time_start = response.data.time_start;
      });
    },
  },
};
