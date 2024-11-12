import axios from 'axios';
import { Author } from './types';

const ENDPOINT = "localhost:8080/api";

export const fetchAuthors = async () => {
  return (await axios.get<Author[]>(ENDPOINT + "/authors")).data;
}

export const fetchAuthorByID = async (id: number) => {
  return (await axios.get<Author>(ENDPOINT + `/authors/${id}`)).data;
}
