<script setup lang="ts">
  import { ref } from 'vue';
  import DataTable from './components/DataTable.vue';
  import { Author } from './types';
  import { fetchAuthorByID } from './services';
</script>

<script lang="ts">
  let authors: Author[] = [
    {id: 0, name: "Brandon Sanderson", email: "sanderson@heartsteel.com", birthday: "19-12-1975"},
    {id: 1, name: "Agatha Christie", email: "agatha@booksclub.com", birthday: "15-9-1890"},
    {id: 2, name: "Stephen King", email: "stephenking@noreply.com", birthday: "21-9-1947"},
    {id: 3, name: "James Patterson", email: "jamespatt@booksclub.com", birthday: "22-4-1947"},
    {id: 4, name: "J. R. R. Tolkien", email: "tolkien@tolkienfoundation.com", birthday: "3-1-1892"},
  ];
  let selected = ref<Author>();
</script>

<template>
  <span v-if="!selected || !selected.books">
    <h1>Authors</h1>
    <DataTable :data="authors" @selected="async (id) => selected = await fetchAuthorByID(id)" />
  </span>
  
  <span v-if="selected && selected.books">
    <h1>Books by {{ selected.name }}</h1>
    <DataTable :data="selected.books"></DataTable>
  </span>
</template>

<style lang="css">
  :root {
    --background: #141414;
    --text: #D8A4A4;
    --hover: #342424;
  }
  html, body {
    margin: 0rem;
    padding: 0rem;
    width: 100%;
    height: 100%;
    font-family: "Monaspace Argon Var";
    font-weight: 200;
    font-size: 16px;
    background-color: var(--background);
    color: var(--text);
    display: flex;
    justify-content: center;
  }
  div#app {
    padding-top: 2rem;
    display: flex;
    flex-direction: column;
    justify-content: flex-start;
    align-items: flex-start;
  }
</style>