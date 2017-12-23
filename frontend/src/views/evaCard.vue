<template>
<Card>
    <div>
        <h3>{{data.course_name}}-{{data.teacher_name}}</h3>
        <Rate show-text :disabled="disabled" v-model="formData.star">
          <span @click="addComment">{{getRandomComment}}</span>
        </Rate>
        <Form ref="eva" :model="formData" :rules="rule" :label-width="0">
            <FormItem  prop="comment">
                <Input type="textarea" :disabled="disabled" v-model="formData.comment" :autosize="{minRows: 5,maxRows: 7}" placeholder="请输入评教内容" />
            </FormItem>
            <FormItem>
                <Button long type="primary" :disabled="disabled" :loading="loading" @click="handleSubmit('eva')">
                    <span v-if="disabled">已经评教</span>
                    <span v-else-if="!loading">评教</span>
                    <span v-else>评教中，请稍候...</span>
                </Button>
            </FormItem>
        </Form>
    </div>
</Card>
</template>

<script>
const comments = {
  1: ["不太负责的老师"],
  2: ["老师还有很大的进步空间"],
  3: ["老师还不错但是需要有提高"],
  4: [
    "老师比较热爱教学，用心经营教学",
    "老师责任心较强",
    "教师具有比较丰富的实践教学经验",
    "教师比较及时反馈作业中出现的问题",
    "该门课程推荐了比较优质的教材"
  ],
  5: [
    "老师热爱教学，用心经营教学",
    "老师责任心强，实践教学目的明确",
    "教师具有丰富的实践教学经验",
    "教师认真考勤，按时检查进度",
    "教师及时反馈作业中出现的问题",
    "该课程推荐了优质教材"
  ]
};
export default {
  name: "eva-card",
  props: {
    data: {
      type: Object,
      require: true
    }
  },
  data() {
    return {
      is_loading: false,
      formData: {
        star: 0,
        comment: ""
      },
      rule: {
        comment: [
          {
            required: true,
            message: "必填哦",
            trigger: "blur"
          },
          {
            type: "string",
            min: 2,
            max: 100,
            message: "字符限制： 2-100",
            trigger: "blur"
          }
        ]
      }
    };
  },
  computed: {
    getRandomComment: function() {
      let star = this.formData.star;
      if (!star || star < 1 || star > 5) return "";
      let len = comments[star].length;
      let rand = Math.round(Math.random() * (len - 1));
      return comments[star][rand];
    },
    loading: function() {
      if ("status" in this.data && this.data.status) {
        return false;
      } else {
        return this.is_loading;
      }
    },
    disabled: function() {
      if ("status" in this.data && this.data.status) {
        this.formData.star = this.data.star;
        this.formData.comment = this.data.comment;
        return true;
      }
    }
  },
  methods: {
    handleSubmit: function(name) {
      this.$refs[name].validate(valid => {
        if (valid && this.formData.star > 0 && this.formData.star < 6) {
          this.is_loading = true;
          let data = this.data;
          data.star = this.formData.star;
          data.comment = this.formData.comment;
          this.$socket.emit("evaluate", data);
        } else {
          this.$Message.error("请先输入正确的信息");
        }
      });
    },

    addComment: function(e) {
      if (e.target.innerText) {
        this.formData.comment = e.target.innerText;
      }
    }
  }
};
</script>

<style>

</style>
