type FilterViteKeys<T> = {
  [K in keyof T]: K extends `VITE_${string}` ? K : never;
}[keyof T];

export type Env<T extends Record<string, string | number>> = {
  [K in FilterViteKeys<T> as K extends `VITE_${infer P}` ? P : never]: T[K];
};
