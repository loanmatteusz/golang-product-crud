export const useAuth = () => {
    const auth = useState<boolean | null>('auth', () => null);
    const config = useRuntimeConfig();

    const register = async (name: string, email: string, password: string) => {
        const { data, error } = await useFetch('/auth/register', {
            method: 'POST',
            body: {
                name,
                email,
                password,
            },
            baseURL:  config.public.apiUrl,
        });

        if (error.value || !data.value) {
            console.error(error);
        }

        return data.value;
    }

    const login = async (email: string, password: string) => {
        const { data, error } = await useFetch<boolean>('/auth/login', {
            method: 'POST',
            body: { email, password },
            baseURL: config.public.apiUrl,
        });

        if (error.value || !data.value) {
            auth.value = false;
            return false;
        }

        auth.value = true;
        return true;
    }

    const logout = () => {
        auth.value = null;
    }

    const isAuthenticated = computed(() => !!auth.value);

    return { auth, register, login, logout, isAuthenticated }
}
