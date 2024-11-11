<script setup lang="ts">
  import { computed } from 'vue';

  const props = defineProps<{ data: Array<Record<string, any>> }>();
  const fields = computed(() => Object.keys(props.data[0]) as Array<keyof typeof props.data[0]>);
</script>

<template>
  <table>
    <thead>
      <th v-for="col in fields" :key="col">{{ col }}</th>
    </thead>
    <tbody>
      <tr v-for="row in data" :key="row.id">
        <td v-for="col in fields" :key="col">{{ row[col] }}</td>
      </tr>
    </tbody>
  </table>
</template>

<style lang="css">
  table {
    border-collapse: collapse;
  }
  thead th {
    padding: .75rem 1.5rem;
    text-align: left;
    text-transform: uppercase;
    font-weight: bold;
  }
  tbody tr {
    text-align: left;
  }
  tbody tr td {
    padding: .75rem 1.5rem;
  }
  tbody tr:hover {
    background-color: var(--hover);
  }
</style>