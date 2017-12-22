<template>
    <div id="login">
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
</template>
<script>
export default {
  data() {
    const validateStudentId = (rule, value, callback) => {
      if (value === "" || isNaN(value)) {
        callback(new Error("您的没有输入学号或者您的学号不是数字！"));
      }else{
          callback();
      }
    };
    const validatePassWord = (rule, value, callback) => {
      if (value === "") {
        callback(new Error("密码不能为空！"));
      }else{
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
      loading: false
    };
  },
  sockets: {
    connect: function() {
      console.log("连接成功！");
    },
    login: function(val) {
      if (val.status!=0){
          this.$Message.error(val.msg);
      }else{
          this.$Message.success("登录成功！")
          //跳转页面
      }
      this.loading = false
    }
  },
  methods: {
    handleSubmit(name) {
      this.$refs[name].validate(valid => {
        console.log("xx",valid);
        if (valid) {
          this.loading = true
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
#login {
  margin-top: 10em;
  width: 100%;
  height: 100%;
}

#login h3 {
  color: #333;
  font-size: 2.5em;
  text-align: center;
  margin-bottom: 1em;
}
</style>
