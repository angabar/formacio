import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query";

export const api = createApi({
    reducerPath: "api",
    baseQuery: fetchBaseQuery({
        baseUrl: "https://jsonplaceholder.typicode.com/",
    }),
    endpoints: (builder) => ({
        getUser: builder.query({
            query: (id) => `users/${id}`,
        }),
    }),
});
