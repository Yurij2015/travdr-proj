<template>
  <div class="row">
    <div class="col-md-12 mt-3">
      <form @submit.prevent="submit">
        <h1 class="h3 mb-3 fw-normal">Редактировать информацию клиента</h1>
        <div class="mb-3">
          <label for="first_name">ФИО</label>
          <input name="first_name" id="first_name" class="form-control" placeholder="ФИО" v-model="data.full_name">
        </div>
        <div class="mb-3">
          <label>Пол</label>
          <select name="role_id" class="form-control" required v-model="data.gender">
            <option disabled selected>{{ data.gender }}</option>
            <option value="Мужской">Мужской</option>
            <option value="Женский">Женский</option>
          </select>
        </div>
        <div class="mb-3">
          <label for="birth_date">Дата рождения: {{ data.birth_date }}</label>
          <input name="birth_date" id="birth_date" type="date" class="form-control"
                 v-model="data.birth_date">
        </div>
        <div class="mb-3">
          <label for="birth_place">Место рождения</label>
          <input name="birth_place" type="text" id="birth_place" class="form-control" placeholder="Место рождения"
                 v-model="data.birth_place">
        </div>
        <div class="mb-3">
          <label for="image">Фотография клиента</label>
          <div class="input-group">
            <input name="image" class="form-control" id="image" placeholder="Изображение" v-model="data.image">
            <ImageUpload @uploaded="data.image = $event"/>
          </div>
        </div>
        <button class="w-100 btn btn-lg btn-primary">Сохранить</button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import {useRoute, useRouter} from "vue-router";
import axios from "axios";
import {onMounted, reactive} from "vue";
import ImageUpload from "@/components/ImageUpload.vue";

export default {
  name: "CustomerEdit",
  components: {ImageUpload},
  setup() {
    const router = useRouter();
    const route = useRoute();
    const data = reactive({
      full_name: '',
      gender: '',
      birth_date: '',
      birth_place: '',
      image: '',
    });

    onMounted(async () => {
      const response = await axios.get(`customers/${route.params.id}`);

      data.full_name = response.data.full_name;
      data.gender = response.data.gender;
      data.birth_date = response.data.birth_date;
      data.birth_place = response.data.birth_place;
      data.image = response.data.image;

    });

    const submit = async () => {
      await axios.put(`customers/${route.params.id}`, data)
      await router.push('/customers')
    }


    return {
      data,
      submit
    }
  }
}
</script>
