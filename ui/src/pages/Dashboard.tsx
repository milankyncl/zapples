import React, { FC } from 'react';
import useSWR from 'swr'
import { client } from '../api/client'

const fetcher = (url: string) => client.get(url).then(res => res.data)

interface FeatureToggle {
    id: number;
    key: string;
    enabled: boolean;
}

interface FeatureTogglesResponse {
    data: FeatureToggle[];
}

export const Dashboard: FC = () => {
    const { data, error, isLoading } = useSWR<FeatureTogglesResponse>('/api/features', fetcher)

    if (error) return <div>failed to load</div>
    if (isLoading || !data) return <div>loading...</div>

    const onToggleCheck = (id: number, enabled: boolean) => async () => {
        await client.put(`/api/features/${id}/toggle`, {
            enabled,
        });
    }

    return <>
        <h1 className="font-medium leading-tight text-2xl mt-0 mb-2 text-black">Features</h1>
        {data.data.map((toggle) => (
            <div className="mb-2 flex">
                <div className="flex-1">
                    <span className="font-medium">{toggle.key}</span>
                </div>
                <div>
                    <label className="inline-flex relative items-center cursor-pointer">
                        <input type="checkbox" className="sr-only peer" defaultChecked={toggle.enabled} onChange={onToggleCheck(toggle.id, !toggle.enabled)} />
                        <div className="w-11 h-6 bg-gray-200 rounded-full peer dark:bg-gray-700 peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-[2px] after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all dark:border-gray-600 peer-checked:bg-blue-600" />
                    </label>
                </div>
            </div>
        ))}
    </>
}