import { configureStore, createSlice } from "@reduxjs/toolkit";

import { api } from "../api/api";

const counterSlice = createSlice({
    name: "counter",
    initialState: {
        value: 0,
    },
    reducers: {
        increment(state) {
            state.value++;
        },
        decrement(state) {
            state.value--;
        },
    },
});

export const { increment, decrement } = counterSlice.actions;

export const store = configureStore({
    reducer: {
        counter: counterSlice.reducer,
        [api.reducerPath]: api.reducer,
    },
    middleware: (getDefaultMiddleware) =>
        getDefaultMiddleware().concat(api.middleware),
});
