<template>
  <div>
    <v-card elevation="2" class="chat">
      <v-card-title>
        我的第一个QQ群 <v-icon :color="color" x-large>mdi-message-text</v-icon>
      </v-card-title>
      <v-divider class="mx-4"></v-divider>

      <v-card-text>
        <div class="chat-text">
          <ul>
            <li v-for="(item, k) in lists" :key="k" :class="item.style">
              {{ item.from }}:{{ item.content }}
            </li>
          </ul>
        </div>
      </v-card-text>

      <v-divider class="mx-4"></v-divider>

      <v-container class="footer">
        <v-row>
          <v-col cols="10">
            <v-text-field label="说点什么吧" v-model="sendTxt"></v-text-field>
          </v-col>
          <v-col>
            <v-btn elevation="2" color="primary" @click="send">发送</v-btn>
          </v-col>
        </v-row>
      </v-container>
    </v-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      lists: [],
      sendTxt: "",
      wsConn: "",
      window: window,
      color: "red",
    };
  },
  mounted() {
    //   console.log(this.window.renderData);
    this.init(this.window.renderData);
  },
  methods: {
    init(renderData) {
      this.wsConn = new WebSocket(
        renderData.Host +
          "?nickname=" +
          renderData.Nickname +
          "&room_id=" +
          renderData.RoomId +
          "&user_id=" +
          renderData.UserId
      );

      // 监听socket连接
      this.wsConn.onopen = this.open;
      // 监听socket错误信息
      this.wsConn.onerror = this.error;
      // 监听socket消息
      this.wsConn.onmessage = this.getMessage;
    },
    open() {
      this.color = "green";
      console.log("socket连接成功");
    },
    error() {
      this.color = "red";
      console.log("连接错误");
    },
    getMessage(msg) {
      let response = JSON.parse(msg.data);

      let style = "left";
      if (response.from != this.window.renderData.Nickname) style = "right";

      response.style = style;
      this.lists.push(response);
    },
    close() {
      this.color = "red";
      console.log("socket已经关闭");
    },
    send() {
      this.wsConn.send(
        JSON.stringify({
          from: this.window.renderData.Nickname,
          content: this.sendTxt,
        })
      );

      this.sendTxt = "";
    },
  },
};
</script>

<style lang="scss">
.chat {
  min-height: 60rem;
  .chat-text {
    min-height: 45rem;
    li {
      list-style: none;
    }
    .left {
      text-align: left;
    }
    .right {
      text-align: right;
    }
  }
}

.text-box {
  width: auto;
}
</style>