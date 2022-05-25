import axios from "axios";

export default {
  namespaced: true,
  state: {
    id: null,
    name: null,
    groups: [],
    default_standard_id: null,
    avatar: null,
  },
  mutations: {
    login() {
      const url = new URL(process.env.VUE_APP_ADMIN_URI + "/oauth/authorize");
      url.searchParams.set("client_id", process.env.VUE_APP_Client_ID);
      url.searchParams.set("redirect_uri", process.env.VUE_APP_REDIRECTURI);
      document.location.assign(url.href);
    },
  },
  actions: {
    info({ commit, state }) {
      if (state.id) return;
      return axios
        .get("/api/admin/info")
        .then((response) => {
          state.name = response.data.name;
          state.groups = response.data.groups.map((groups) => groups.name);
          state.default_standard_id = response.data.standard_id;
          state.avatar = response.data.avatar;
        })
        .catch((error) => {
          if (error.response.status == 401) {
            commit("logout");
            commit("login");
          }
        });
    },

    logout({ state }) {
      state.id = null;
      state.name = null;
      state.group = [];
      state.default_standard_id = null;
      state.avatar = null;
      return axios.post("/api/admin/logout");
    },

    login({ state }, code) {
      return axios.post("/api/admin/login/" + code).then((response) => {
        state.name = response.data.name;
        state.groups = response.data.groups.map((groups) => groups.name);
        state.default_standard_id = response.data.default_standard_id;
        state.avatar = response.data.avatar;
      });
    },

    setDefault({ state }, standard_id) {
      return axios({
        method: "post",
        url: "/api/admin/set-default-standard",
        data: new URLSearchParams({ standard_id: standard_id }),
      })
        .then(() => {
          state.default_standard_id = standard_id;
        })
        .finally(() => (this.submit_loading = false));
    },
  },
  getters: {},
};
