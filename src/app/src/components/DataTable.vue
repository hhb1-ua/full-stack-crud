<script setup lang="ts">
  import { computed, defineEmits } from 'vue';

  const props = defineProps<{ data: Array<Record<string, any>> }>();
  const fields = computed(() => Object.keys(props.data[0]) as Array<keyof typeof props.data[0]>);
  const emit = defineEmits<{ (e: 'selected', data: number): void }>();
</script>

<template>
  <table>
    <thead>
      <th v-for="col in fields" :key="col">{{ col }}</th>
    </thead>
    <tbody>
      <tr v-for="row in data" :key="row.id" @click="emit('selected', row.id)">
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
    background-color: var(--text);
    color: var(--background);
  }
  tbody tr {
    text-align: left;
  }
  tbody tr td {
    padding: .75rem 1.5rem;
    border-top: 1px solid var(--text);
  }
  tbody tr:hover {
    background-color: var(--hover);
  }
</style>