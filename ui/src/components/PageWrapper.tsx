import React, { FC } from 'react';
import { Logo } from "./atoms/Logo";

interface Props {
    children: React.ReactNode;
}

export const PageWrapper: FC<Props> = ({ children }) => {
    return <>
        <div className="text-center mb-4">
            <Logo />
        </div>
        <div className="relative bg-white px-6 py-6 shadow-xl ring-1 ring-gray-900/5 sm:mx-auto sm:max-w-lg sm:rounded-lg sm:px-10">
            <div className="mx-auto max-w-md">
                <div className="divide-y divide-gray-300/50">
                    <div className="space-y-6 py-3 text-base leading-7 text-gray-600">
                        {children}
                    </div>
                    {/*
                    <div className="pt-8 text-base font-semibold leading-7">
                        <p className="text-gray-900">Want to dig deeper into Tailwind?</p>
                        <p>
                            <a href="https://tailwindcss.com/docs" className="text-sky-500 hover:text-sky-600">Read the
                                docs &rarr;</a>
                        </p>
                    </div>
                    */}
                </div>
            </div>
        </div>
    </>
}
