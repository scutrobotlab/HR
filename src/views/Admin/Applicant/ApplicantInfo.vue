<!-- 报名表详情页面 -->
<template>
  <v-main>
    <ErrorAlert critical v-model="error" />
    <v-list tile>
      <v-list-item>
        <v-list-item-avatar>
          <v-icon>mdi-account</v-icon>
        </v-list-item-avatar>
        <v-list-item-content>
          <v-list-item-subtitle>{{ $t("name") }}</v-list-item-subtitle>
          <v-list-item-title>{{ applicant.name }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>

      <v-row align="center" class="mr-0">
        <v-col class="grow py-0 pr-0">
          <v-list-item :href="'tel:' + applicant.phone">
            <v-list-item-avatar>
              <v-icon color="primary">mdi-phone</v-icon>
            </v-list-item-avatar>
            <v-list-item-content>
              <v-list-item-subtitle>{{ $t("phone") }}</v-list-item-subtitle>
              <v-list-item-title>{{ applicant.phone }}</v-list-item-title>
            </v-list-item-content>
          </v-list-item>
        </v-col>
        <v-col class="shrink py-0 px-3">
          <v-btn
            icon
            large
            :href="
              'sms:' +
              applicant.phone +
              ';?&body=【华工机器人实验室】' +
              applicant.name +
              '同学你好，欢迎你报名参加机器人实验室面试。\
我们看到你尚未确定面试时间，请看到短信后在报名系统里选择好你的时间，以便正常参加面试。\
收到请回复“' +
              applicant.name +
              '收到”'
            "
          >
            <v-icon>mdi-message-text</v-icon>
          </v-btn>
        </v-col>
      </v-row>
      <v-list-item>
        <v-list-item-avatar>
          <v-icon>mdi-gender-{{ applicant.gender }}</v-icon>
        </v-list-item-avatar>
        <v-list-item-content>
          <v-list-item-subtitle>{{ $t("gender") }}</v-list-item-subtitle>
          <v-list-item-title>{{ $t(applicant.gender) }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-list-item>
        <v-list-item-avatar>
          <v-icon>mdi-heart</v-icon>
        </v-list-item-avatar>
        <v-list-item-content>
          <v-list-item-subtitle>{{ $t("intent") }}</v-list-item-subtitle>
          <div>
            <GroupChip
              v-for="intent in applicant.intents"
              :key="intent.key"
              :intent="intent"
            ></GroupChip>
          </div>
        </v-list-item-content>
      </v-list-item>
      <v-list-item v-for="admit in applicant.admits" :key="admit.key">
        <v-list-item-avatar>
          <v-icon>mdi-marker-check</v-icon>
        </v-list-item-avatar>
        <v-list-item-content>
          <v-list-item-subtitle
            >{{ admit.admin.name }}于{{ format(admit.updated_at) }}录取</v-list-item-subtitle
          >
          <v-list-item-title>{{ admit.group }}</v-list-item-title>
        </v-list-item-content>
        <v-list-item-action v-if="admin.groups.includes(admit.group)">
          <CancleAdmit :admit="admit" :applicant_name="applicant.name" />
        </v-list-item-action>
      </v-list-item>
      <v-list-item two-line v-for="(item, key) in applicant.form" :key="key">
        <v-list-item-avatar />
        <v-list-item-content>
          <v-list-item-subtitle>{{ form(key) }}</v-list-item-subtitle>
          <v-list-item-title>{{ item }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-list-item>
        <v-list-item-avatar />
        <v-list-item-content>
          <v-list-item-subtitle>
            {{ format(applicant.created_at) }}报名
            <span v-if="applicant.created_at !== applicant.updated_at"
              >，{{ format(applicant.updated_at) }}修改过</span
            >。
          </v-list-item-subtitle>
        </v-list-item-content>
      </v-list-item>
    </v-list>
    <div class="d-block text-center pa-10">
      <v-dialog v-model="dialog" width="500">
        <template v-slot:activator="{ on, attrs }">
          <v-btn outlined color="error" v-bind="attrs" v-on="on" :loading="delete_loading"
            ><v-icon left>mdi-delete</v-icon>删除</v-btn
          >
        </template>
        <v-card>
          <v-card-title> 警告：请仔细阅读 </v-card-title>
          <v-card-text>
            优先推荐给面试者打低分，使之不出现在列表中。<br />
            删除操作，将会从系统中，所有组别，删除此面试者。<br />
            即便此面试者已面试过其他组，所有组面试官也都不能再看到这个面试者。<br />
            请谨慎使用删除功能。<br />
            如果误操作，可以联系网管进入数据库进行恢复。<br />
          </v-card-text>
          <v-card-actions>
            <v-btn color="primary" text @click="dialog = false">取消</v-btn>
            <v-spacer></v-spacer>
            <v-btn color="error" text @click="del">确认删除</v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
    </div>
  </v-main>
</template>

<script>
import moment from "moment";
import axios from "axios";
import GroupChip from "@/components/GroupChip.vue";
import errorMixin from "@/mixins/errorMixin.js";
import CancleAdmit from "@/components/CancleAdmit.vue";

export default {
  mixins: [errorMixin],
  components: { GroupChip, CancleAdmit },
  data: () => ({
    drawer: null,
    delete_loading: false,
    cancel_loading: false,
    dialog: false,
  }),
  computed: {
    groups_loaded() {
      return this.$store.state.groups.loaded;
    },
    groups() {
      return this.$store.state.groups.list;
    },
    applicant() {
      return this.$store.state.adminApplicant.applicant;
    },
    admin() {
      return this.$store.state.admin;
    },
  },
  methods: {
    format(dateTime) {
      return moment(dateTime).fromNow();
    },
    form(key) {
      var form = this.$store.state.form.form;
      if (!form) return this.$t(key);
      else {
        var feild = form.find((item) => item.key === key);
        if (!feild) return this.$t(key);
        return feild.name;
      }
    },
    del() {
      this.delete_loading = true;
      this.errorHandler(
        axios({
          method: "delete",
          url: "/api/applicant/" + this.$route.params.appId,
        })
      )
        .then(() => this.$router.go("-1"))
        .finally(() => (this.delete_loading = false));
    },
  },
  mounted() {
    this.errorHandler(this.$store.dispatch("form/get"));
    this.$store.dispatch("groups/get");
  },
};
</script>

<style scoped>
.v-list-item__title {
  white-space: normal;
  height: fit-content;
}
</style>
