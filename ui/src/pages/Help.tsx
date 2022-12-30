import React, { FC } from 'react';

export const Help: FC = () => {
    return <>
        <p>An advanced online playground for Tailwind CSS, including support for things like:</p>
        <ul className="space-y-4">
            <li className="flex items-center">
                <svg className="h-6 w-6 flex-none fill-sky-100 stroke-sky-500 stroke-2"
                     stroke-linecap="round" stroke-linejoin="round">
                    <circle cx="12" cy="12" r="11"/>
                    <path d="m8 13 2.165 2.165a1 1 0 0 0 1.521-.126L16 9" fill="none"/>
                </svg>
                <p className="ml-4">
                    Customizing your
                    <code className="text-sm font-bold text-gray-900">tailwind.config.js</code> file
                </p>
            </li>
            <li className="flex items-center">
                <svg className="h-6 w-6 flex-none fill-sky-100 stroke-sky-500 stroke-2"
                     stroke-linecap="round" stroke-linejoin="round">
                    <circle cx="12" cy="12" r="11"/>
                    <path d="m8 13 2.165 2.165a1 1 0 0 0 1.521-.126L16 9" fill="none"/>
                </svg>
                <p className="ml-4">
                    Extracting classes with
                    <code className="text-sm font-bold text-gray-900">@apply</code>
                </p>
            </li>
            <li className="flex items-center">
                <svg className="h-6 w-6 flex-none fill-sky-100 stroke-sky-500 stroke-2"
                     stroke-linecap="round" stroke-linejoin="round">
                    <circle cx="12" cy="12" r="11"/>
                    <path d="m8 13 2.165 2.165a1 1 0 0 0 1.521-.126L16 9" fill="none"/>
                </svg>
                <p className="ml-4">Code completion with instant preview</p>
            </li>
        </ul>
        <p>Perfect for learning how the framework works, prototyping a new idea, or creating a demo to
            share online.</p>
    </>
}