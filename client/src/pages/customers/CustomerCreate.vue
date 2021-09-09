<template>
  <div class="row">
    <div class="col-md-12 mt-3">
      <form @submit.prevent="submit">
        <h1 class="h3 mb-3 fw-normal">Добавить клиента</h1>
        <div class="mb-3">
          <label for="first_name">ФИО</label>
          <input name="first_name" id="first_name" class="form-control" placeholder="ФИО" v-model="data.full_name">
        </div>
        <div class="mb-3">
          <label>Пол</label>
          <select name="role_id" class="form-control" required v-model="data.gender">
            <option disabled selected>Выберите необходимое значение</option>
            <option value="Мужской">Мужской</option>
            <option value="Женский">Женский</option>
          </select>
        </div>
        <div class="mb-3">
          <label for="last_name">Дата рождения</label>
          <input name="last_name" id="last_name" type="date" class="form-control" placeholder="Фамилия"
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
import {useRouter} from "vue-router";
import axios from "axios";
import {reactive} from "vue";
import ImageUpload from "@/components/ImageUpload.vue";

export default {
  name: "CustomerCreate",
  components: {ImageUpload},
  setup() {
    const router = useRouter();
    const data = reactive({
      full_name: '',
      gender: '',
      birth_date: '',
      birth_place: '',
      image: '',
    });

    const submit = async () => {
      await axios.post('customers', data)
      await router.push('/customers')
    }


    return {
      data,
      submit
    }
  }
}
</script>
