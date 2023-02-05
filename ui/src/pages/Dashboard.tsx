import React, {FC} from 'react';
import useSWR from 'swr'
import {client} from '../api/client'
import {Button, ButtonSize} from "../components/atoms/Button";
import {PlusIcon} from "@heroicons/react/24/solid";
import {useNavigate} from "react-router-dom";
import {PageHeading} from "../components/atoms/PageHeading";
import {FeatureItem} from "./features/FeatureItem";
import {Feature} from "../api/models/feature";
import {ErrorMessage} from "../components/atoms/ErrorMessage";
import {LoadingMessage} from "../components/atoms/LoadingMessage";

const fetcher = (url: string) => client.get(url).then(res => res.data)

interface FeatureTogglesResponse {
    data: Feature[];
}

function noItemsYet() {
    return <>
        <p className="text-sm text-gray-500 text-center italic py-6">
            No features created yet
        </p>
    </>;
}

export const Dashboard: FC = () => {
    const navigate = useNavigate();
    const { data, error, isLoading, mutate } = useSWR<FeatureTogglesResponse>('/features', fetcher);

    if (error) {
        return <ErrorMessage>Failed to load features</ErrorMessage>
    }
    if (isLoading || !data) {
        return <LoadingMessage />
    }

    const createNewFeature = () => navigate('/features/create');

    const handleToggle = (id: number) => async (enabled: boolean) => {
        await client.put(`/features/${id}/toggle`, {
            enabled,
        });
    }
    const handleRemove = (id: number) => async () => {
        if (window.confirm('Are you sure?')) {
            await client.delete(`/features/${id}`);
            await mutate();
        }
    }
    const handleEdit = (id: number) => () => navigate(`/features/${id}`);

    return <>
        <div className="flex items-center mb-4">
            <PageHeading>Features</PageHeading>
            <div className="text-right flex-1">
                <Button size={ButtonSize.Sm} onClick={createNewFeature} className="inline-flex items-center">
                    <PlusIcon className="mr-2 -ml-1 w-4 h-4 text-white" />
                    <span>Create new</span>
                </Button>
            </div>
        </div>
        <div className="divide-y divide-gray-300/50">
            {data.data.length === 0 && noItemsYet()}
            {data.data.map((feature) => <FeatureItem
                feature={feature}
                onToggle={handleToggle(feature.id)}
                onRemove={handleRemove(feature.id)}
                onEdit={handleEdit(feature.id)}
            />)}
        </div>
    </>
}
