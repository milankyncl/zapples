import React, { FC } from 'react';
import useSWR from 'swr'
import axios from 'axios'

const fetcher = (url: string) => axios.get(url).then(res => res.data)

interface FeatureToggle {
    name: string;
    enabled: boolean;
}

interface FeatureTogglesResponse {
    data: FeatureToggle[];
}

export const Dashboard: FC = () => {
    const { data, error, isLoading } = useSWR<FeatureTogglesResponse>('/api/admin/feature-storage', fetcher)

    if (error) return <div>failed to load</div>
    if (isLoading || !data) return <div>loading...</div>

    return <>
        {data.data.map((toggle) => (
            <div>
                {toggle.name}
            </div>
        ))}
    </>
}