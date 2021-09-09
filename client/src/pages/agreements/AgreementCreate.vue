<template>
  <div class="row">
    <div class="col-md-12 mt-3">
      <form @submit.prevent="submit">
        <h1 class="h3 mb-3 fw-normal">Добавить соглашение</h1>
        <div class="mb-3">
          <label for="agreement_name">Номер соглашения</label>
          <input name="agreement_name" id="agreement_name" class="form-control" placeholder="Номер соглашения"
                 required v-model="data.agreement_name">
        </div>
        <div class="mb-3">
          <label for="description">Описание соглашения</label>
          <textarea name="description" id="description" class="form-control" placeholder="Описание соглашения"
                    required v-model="data.description"></textarea>
        </div>
        <div class="mb-3">
          <label>Содержимое соглашения</label>
          <editor v-model="data.content" initial-value="{{data.content}}" id="content" apiKey="7pxxebfatsrkizfz23o8eh0fz5wpqja4k03eq2z1hzpyqy5h"/>
        </div>
        <button class="w-49 btn btn-lg btn-warning">Сохранить соглашение</button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import {useRouter} from "vue-router";
import axios from "axios";
import {onMounted, reactive, ref} from "vue";
import Editor from '@tinymce/tinymce-vue'

export default {
  name: "AgreementCreate",
  components: {
    'editor': Editor
  },
  setup() {
    const router = useRouter();
    const data = reactive({
      agreement_name: '',
      description: '',
      content: '',
    });

    const submit = async () => {
      await axios.post('agreements', data)
      await router.push('/agreements')
    }

    return {
      data,
      submit
    }
  }
}
</script>
