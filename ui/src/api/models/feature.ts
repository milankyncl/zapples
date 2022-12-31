export interface Feature {
    id: number;
    key: string;
    description: string | null;
    enabled: boolean;
    createdAt: Date;
}