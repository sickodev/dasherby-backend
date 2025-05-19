package supabase

import (
	"log"

	"github.com/sickodev/dasherby-backend/internal/config"
	"github.com/supabase-community/supabase-go"
)

var Client *supabase.Client

func Init(cfg *config.Config){
	var err error
	Client, err = supabase.NewClient(cfg.SupabaseURL, cfg.SupabaseKey, nil)

	if err != nil {
		log.Fatalf("Failed to intialize Supabase client: %v", err)
	}
}