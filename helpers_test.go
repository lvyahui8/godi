package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type TS01 struct {
	Bean
}

type TS02 struct {
	*Bean
}

type TS03 struct {
}

func (T TS03) BeanName() string {
	panic("implement me")
}

func (T TS03) Init() error {
	panic("implement me")
}

type TS04 struct {
}

func (T *TS04) BeanName() string {
	panic("implement me")
}

func (T *TS04) Init() error {
	panic("implement me")
}

func TestIsBean(t *testing.T) {
	require.True(t, IsBean(&TS01{}))
	require.True(t, IsBean(TS01{}))
	require.True(t, IsBean(&TS02{}))
	require.True(t, IsBean(&TS02{}))
	require.True(t, IsBean(TS03{}))
	require.True(t, IsBean(&TS03{}))
	require.True(t, IsBean(TS04{}))
	require.True(t, IsBean(&TS04{}))
}
