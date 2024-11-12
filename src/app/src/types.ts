export type Author = {
  id: number;
  name: string;
  email: string;
  birthday: string;
  books?: Array<Book>;
};

export type Book = {
  id: number;
  title: string;
  category: 'Unknown' | 'Fantasy' | 'Fiction' | 'Romance' | 'Horror' | 'Thriller' | 'Mistery';
};
