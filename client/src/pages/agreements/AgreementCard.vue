<template>
  <div class="row">
    <div class="col-md-12 mt-3">
      <form @submit.prevent="submit">
        <h1 class="h3 mb-3 fw-normal">Карточка агента - {{ data.full_name }} </h1>
        <div class="mb-3">
          <label for="first_name">ФИО</label>
          <input name="first_name" id="first_name" class="form-control" placeholder="ФИО" v-model="data.full_name">
        </div>
        <div class="mb-3">
          <label for="phone_number">Номер телефона</label>
          <input name="phone_number" id="phone_number" type="tel" class="form-control" placeholder="Номер телефона"
                 v-model="data.phone_number">
        </div>
        <div class="mb-3">
          <label for="organization_id">Организация</label>
          <select v-model="data.organization_id" id="organization_id" name="organization_id" class="form-control">
            <option v-for="organization of organizations" :key="organization.id" :value="organization.id">
              {{ organization.org_name }}
            </option>
          </select>
        </div>
        <div class="mb-3">
          <label for="image">Фотография агента</label>
          <div class="input-group">
            <img :src="data.image" width="150" alt=""/>
            <input name="image" class="form-control" id="image" placeholder="Изображение" v-model="data.image">
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
import {onMounted, reactive, ref} from "vue";

export default {
  name: "AgreementCard",
  setup() {
    const router = useRouter();
    const route = useRoute();
    const data = reactive({
      full_name: '',
      phone_number: '',
      organization_id: '',
      image: '',
    });

    const organizations = ref([])
    onMounted(async () => {
      const organizationResponse = await axios.get('organizations')
      organizations.value = organizationResponse.data.data;

      const response = await axios.get(`agents/${route.params.id}`);

      data.full_name = response.data.full_name;
      data.phone_number = response.data.phone_number;
      data.organization_id = response.data.organization_id;
      data.image = response.data.image;
    });

    const submit = async () => {
      await axios.put(`agents/${route.params.id}`, data)
      await router.push('/agents')
    }

    return {
      data,
      submit,
      organizations
    }
  }
}
</script>
