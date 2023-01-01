import React, {FC, useEffect, useState} from 'react';
import {useNavigate, useParams} from "react-router-dom";
import {BackButton} from "../../components/atoms/BackButton";
import {PageHeading} from "../../components/atoms/PageHeading";
import {FeatureForm, FeatureFormDto} from "../../components/features/FeatureForm";
import {client} from "../../api/client";
import useSWRImmutable from "swr/immutable";
import {Feature} from "../../api/models/feature";

const fetcher = (url: string) => client.get(url).then(res => res.data)

interface FeatureResponse {
    data: Feature;
}

export const EditFeature: FC = () => {
    const navigate = useNavigate();
    const params = useParams();
    const { data, error, isLoading, mutate } = useSWRImmutable<FeatureResponse>(`/features/${params.id}`, fetcher, {
        revalidateOnMount: true,
        revalidateOnFocus: true,
        keepPreviousData: false,
    })
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        mutate()
        console.log('x');
    }, [params.id]);

    const onSubmitForm = (id: number) => async (data: FeatureFormDto) => {
        setLoading(true);
        try {
            await client.put(`/features/${id}`, {
                key: data.key,
                description: data.description,
            })
        } catch (e) {
            console.error(e, 'Error happened while updating feature')
        } finally {
            setLoading(false);
            navigate('/');
        }
    };

    if (error) return <div>failed to load</div>
    if (isLoading || !data) return <div>loading...</div>

    return <>
        <div className="flex items-center mb-2">
            <BackButton onClick={() => navigate('/')}></BackButton>
            <PageHeading>Edit feature</PageHeading>
        </div>
        <FeatureForm initialData={data.data} loading={loading} onSubmit={onSubmitForm(data.data.id)} />
    </>
}