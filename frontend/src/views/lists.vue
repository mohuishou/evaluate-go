<template>
    <Row>
        <Col v-for="n in 10" :key="n" :xs="{ span: 24}" :md="{ span: 8}" :lg="{ span: 4}">
         <eva-card :data="data"></eva-card>
        </Col>
    </Row>
          
</template>
<script>
import evaCard from "./evaCard.vue";
export default {
components:{
    "eva-card": evaCard
},
  data() {
    return {
      value1: 0,
      start: 0,
      data: {
        course_name: "test",
        teacher_name: "test2",
        star: 5,
        comment: "test",
        status: 1
      }
    };
  }
};
</script>
<style scoped>
.card {
  margin: 5px !important;
}
</style>

