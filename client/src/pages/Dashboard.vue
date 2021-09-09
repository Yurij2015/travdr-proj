<template>
  <h2>Отчет по продажам туристического агенства</h2>
  <div style="height: 600px" id="chart">

  </div>
</template>

<script lang="ts">
import {onMounted} from "vue";
import * as c3 from 'c3';

export default {
  name: "Dashboard",
  setup() {
    onMounted(async () => {
      const chart = c3.generate({
        bindto: '#chart',
        data: {
          x: 'Дата',
          columns: [
            ['Дата', '2021-1-1', '2021-1-2', '2021-1-3', '2021-1-4', '2021-1-5', '2021-1-6'],
            ['Показатель', 100, 200, 300, 150, 320, 280]
          ],
          types: {
            Показатель: 'bar'
          }
        },
        axis: {
          x: {
            type: 'timeseries',
            tick: {
              format: '%Y-%m-%d'
            }
          }
        }
      });
    });
  }
}
</script>
