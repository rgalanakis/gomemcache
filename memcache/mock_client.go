package memcache


type MockClient struct {
	Items  map[string]string
	Errors map[string]error
}

func (mc *MockClient) DeleteAll() error {
	mc.Items = make(map[string]string)
	mc.Errors = make(map[string]error)
}

func (mc *MockClient) Replace(item *Item) error {
	_, err := mc.get(item.Key)
	if err != nil {
		return err
	}
	mc.Items[item.Key] = string(item.Value)
	return nil
}

func (mc *MockClient) Increment(key string, delta uint64) (newValue uint64, err error) {
	panic("implement me")
}

func (mc *MockClient) Decrement(key string, delta uint64) (newValue uint64, err error) {
	panic("implement me")
}

func NewMockClient() *MockClient {
	return &MockClient{Items: make(map[string]string)}
}

func (mc *MockClient) Add(item *Item) error {
	v, err := mc.get(item.Key)
	if err != nil {
		return err
	}
	if v != "" {
		return ErrNotStored
	}
	mc.Items[item.Key] = string(item.Value)
	return nil
}

func (mc *MockClient) Delete(key string) error {
	v, err := mc.get(key)
	if err != nil {
		return err
	}
	if v == "" {
		return ErrCacheMiss
	}
	return nil
}

func (mc *MockClient) Get(key string) (*Item, error) {
	v, err := mc.get(key)
	if err != nil {
		return nil, err
	}
	if v == "" {
		return nil, ErrCacheMiss
	}
	return NewItem(key, []byte(v), 0), nil
}

func (mc *MockClient) Set(item *Item) error {
	if err := mc.Errors[item.Key]; err != nil {
		return err
	}
	mc.Items[item.Key] = string(item.Value)
	return nil
}

func (mc *MockClient) get(key string) (string, error) {
	return mc.Items[key], mc.Errors[key]
}

