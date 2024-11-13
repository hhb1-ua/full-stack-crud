import axios from 'axios';
import { Author, Message } from './types';
import { z } from 'zod';

const ENDPOINT = 'localhost:8080/api';

const authorSchema = z.object({
  id: z.number().positive().optional(),
  name: z.string(),
  email: z.string().email(),
  birthday: z.string().date(),
});

const bookSchema = z.object({
  id: z.number().positive().optional(),
  title: z.string(),
  category: z.string(),
});

// TODO: Parse the responses with Zod, ensuring data integrity
// TODO: Copy-paste the Author services for books

export const fetchAuthors = async () => {
  return (await axios.get<Author[] | Message>(ENDPOINT + '/authors')).data;
}

export const fetchAuthorByID = async (id: number) => {
  return (await axios.get<Author | Message>(ENDPOINT + `/authors/${id}`)).data;
}

export const createAuthor = async (obj: Omit<Author, 'id'>) => {
  authorSchema.parse(obj);
  return (await axios.post<Author | Message>(ENDPOINT + '/authors', obj));
}

export const updateAuthor = async (obj: Author) => {
  authorSchema.parse(obj);
  return (await axios.put<Author | Message>(ENDPOINT + `/authors/${obj.id}`, obj)).data;
}

export const deleteAuthor = async (id: number) => {
  return (await axios.delete<Message>(ENDPOINT + `/authors/${id}`)).data;
}
