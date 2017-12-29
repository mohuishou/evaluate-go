<template>
    <div id="loading" v-if="!isConnected">
      <Spin fix>
        <Icon type="load-c" size=18 class="spin-icon-load"></Icon>
        <div>Loading</div>
      </Spin>
    </div>
    <div v-else-if="!isLogin" id="login">
        <Row id="row">
            <Col :xs="{ span: 22, offset: 1 }" :md="{ span: 12, offset: 6 }" :lg="{ span: 8, offset: 8 }">
                <h3>四川大学快捷评教</h3>
                <Form ref="user" :model="user" :rules="rule" :label-width="40">
                    <FormItem label="学号" prop="student_id">
                        <Input type="text" v-model="user.student_id" />
                    </FormItem>
                    <FormItem label="密码" prop="password">
                        <Input type="password" v-model="user.password" />
                    </FormItem>
                    <FormItem>
                        <Button long type="primary" :loading="loading" @click="handleSubmit('user')">
                            <span v-if="!loading">登录</span>
                            <span v-else>登录中，请稍候...</span>
                        </Button>
                    </FormItem>
                </Form>
            </Col>
        </Row>
    </div>
    <div id="lists" v-else>
    <h3>四川大学快捷评教</h3>
    <Row>
        <Col v-for="(val,k) in getEvaArr" :key="k" :xs="{ span: 24}" :md="{ span: 8}" :lg="{ span: 6}">
         <eva-card class="eva-card" :data="val"></eva-card>
        </Col>
    </Row>
    </div>
</template>
<script>
import evaCard from "./evaCard.vue";
export default {
  components: {
    "eva-card": evaCard
  },
  data() {
    const validateStudentId = (rule, value, callback) => {
      if (value === "" || isNaN(value)) {
        callback(new Error("您的没有输入学号或者您的学号不是数字！"));
      } else {
        callback();
      }
    };
    const validatePassWord = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("密码不能为空！"));
      } else {
        callback();
      }
    };

    return {
      user: {
        password: "",
        student_id: ""
      },
      rule: {
        student_id: [{ validator: validateStudentId, trigger: "blur" }],
        password: [{ validator: validatePassWord, trigger: "blur" }]
      },
      isLogin: false,
      loading: false,
      evaArr: [{}],
      is_count: false,
      evaCount: 0,
      isConnected: false
    };
  },
  sockets: {
    connect: function() {
      this.isConnected = true
      console.log("连接成功！");
    },
    connect_error: function(val){
      this.$Message.error("连接服务器失败！");
    },
    login: function(val) {
      if (val.status != 0) {
        this.$Message.error(val.msg);
      } else {
        this.$Message.success("登录成功！");
      }
      this.loading = false;
    },
    getEvaluateList: function(val) {
      if (val.status != 0) {
        this.$Message.error(val.msg);
      } else {
        this.evaArr = val.data;
        this.$Message.success("评教列表获取成功！");
        this.isLogin = true;
      }
    },
    evaluate: function(val) {
      if (val.status != 0) {
        this.$Message.error(val.msg);
      } else {
        this.$Message.success(val.msg);
        for (let i = 0; i < this.evaArr.length; i++) {
          if (
            this.evaArr[i].teacher_id == val.data.teacher_id &&
            this.evaArr[i].course_id == val.data.course_id
          ) {
            this.$set(this.evaArr, i, val.data);
            this.evaCount++;
            if (this.evaCount == this.evaArr.length) {
              this.$Message.success("所有课程您已评教成功！");
            }
            return;
          }
        }
      }
    }
  },
  computed: {
    getEvaArr() {
      if (this.is_count == false) {
        this.evaArr.forEach(e => {
          this.evaCount += e.status;
        });
        this.is_count = true;
      }
      return this.evaArr;
    }
  },
  methods: {
    handleSubmit(name) {
      this.$refs[name].validate(valid => {
        if (valid) {
          this.loading = true;
          this.$socket.emit("login", this.user);
        } else {
          this.$Message.error("请先输入正确的账号密码");
        }
      });
    }
  }
};
</script>

<style scoped>
#loading {
  display: inline-block;
  width: 100%;
  height: 100%;
  position: relative;
}
.spin-icon-load {
  animation: ani-spin 1s linear infinite;
}
@keyframes ani-spin {
  from {
    transform: rotate(0deg);
  }
  50% {
    transform: rotate(180deg);
  }
  to {
    transform: rotate(360deg);
  }
}
#login {
  margin-top: 10em;
  width: 100%;
  height: 100%;
}

h3 {
  color: #333;
  font-size: 2.2em;
  text-align: center;
  margin-top: 0.8em;
  margin-bottom: 0.8em;
}

.eva-card {
  margin: 0.8em;
}
</style>
