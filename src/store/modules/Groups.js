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
      if (state.loaded) return;
      return axios.get("https://my.scutbot.cn/api/group/groups").then((response) => {
        state.list = response.data.groups.map(g => g.name);
        state.loaded = true;
      });
    },
  },
};
