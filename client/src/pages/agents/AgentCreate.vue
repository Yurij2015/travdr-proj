<template>
  <div class="row">
    <div class="col-md-12 mt-3">
      <form @submit.prevent="submit">
        <h1 class="h3 mb-3 fw-normal">Добавить агента</h1>
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
            <input name="image" class="form-control" id="image" placeholder="Изображение" v-model="data.image">
            <AgentPhotoUpload @uploaded="data.image = $event"/>
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
import {onMounted, reactive, ref} from "vue";
import AgentPhotoUpload from "@/components/AgentPhotoUpload.vue";

export default {
  name: "AgentCreate",
  components: {AgentPhotoUpload},
  setup() {
    const router = useRouter();
    const data = reactive({
      full_name: '',
      phone_number: '',
      organization_id: '',
      image: '',
    });

    const organizations = ref([])
    onMounted(async () => {
      const {data} = await axios.get('organizations')
      organizations.value = data.data;
    })

    const submit = async () => {
      await axios.post('agents', data)
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
