import React from 'react'
import ReactDOM from 'react-dom/client'
import {
    createHashRouter,
    RouterProvider,
} from "react-router-dom";

import './index.css';
import {PageWrapper} from "./components/PageWrapper";
import {Dashboard} from "./pages/Dashboard";
import {CreateFeature} from "./pages/features/CreateFeature";

const router = createHashRouter([
    {
        path: "/",
        element: <Dashboard />,
    },
    {
        path: '/features/create',
        element: <CreateFeature />,
    }
]);

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <PageWrapper>
            <RouterProvider router={router} />
        </PageWrapper>
    </React.StrictMode>,
)
